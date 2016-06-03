package main

import (
	"fmt"
	"sync"

	msg "github.com/scootdev/scoot/messages"
	ci "github.com/scootdev/scoot/sched/clusterimplementations"
	cm "github.com/scootdev/scoot/sched/clustermembership"
	distributor "github.com/scootdev/scoot/sched/distributor"
)

/* demo code */
func main() {

	cluster := ci.StaticLocalNodeClusterFactory(10)
	fmt.Println("clusterMembers:", cluster.Members())

	workCh := make(chan msg.Job)
	distributor := &distributor.RoundRobin{}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		generateTasks(workCh, 100)
		wg.Done()
	}()

	go func() {
		scheduleWork(workCh, cluster, distributor)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("\nTerminating Program All Work Completed")
}

func scheduleWork(
	workCh <-chan msg.Job,
	cluster cm.Cluster,
	distributor distributor.Distributor) {

	var wg sync.WaitGroup
	for work := range workCh {
		node := distributor.DistributeWork(work, cluster)

		wg.Add(1)
		go func(w msg.Job, n cm.Node) {
			defer wg.Done()
			//Todo: error handling, what if request fails
			for _, task := range w.Tasks {
				n.SendMessage(task)
			}
		}(work, node)
	}

	wg.Wait()
}

/*
 * Generates work to send on the channel, using
 * Unbuffered channel because we only want to pull
 * more work when we can process it.
 *
 * For now just generates dummy tasks up to numTasks,
 * In reality this will pull off of work queue.
 */
func generateTasks(work chan<- msg.Job, numTasks int) {

	for x := 0; x < numTasks; x++ {

		work <- msg.Job{
			Id:      fmt.Sprintf("Job %d", x),
			Jobtype: "testTask",
			Tasks: []msg.Task{
				msg.Task{
					Id:       fmt.Sprintf("Job %d, Task 1", x),
					Commands: []string{"testcmd", "testcmd2"},
				},
			},
		}
	}
	close(work)
}
