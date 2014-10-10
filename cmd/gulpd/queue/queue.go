package queue

import (
	"github.com/megamsys/libgo/amqp"
	log "code.google.com/p/log4go"
	"github.com/megamsys/gulp/policies"
	"encoding/json"
)

type QueueServer struct {
	ListenAddress string
	chann          chan []byte
	shutdown      chan bool
}
//interface arguments
func NewServer(listenAddress string) *QueueServer {
	log.Info("Create New queue server")
	self := &QueueServer{}

	self.ListenAddress = listenAddress
	self.shutdown = make(chan bool, 1)
     log.Info(self)
	return self
}



func (self *QueueServer) ListenAndServe() {
	factor, err := amqp.Factory()
	if err != nil {
		log.Error("Failed to get the queue instance: %s", err)
	}
	
	pubsub, err := factor.Get(self.ListenAddress)
	if err != nil {
		log.Error("Failed to get the queue instance: %s", err)
	}
	
	res := &policies.Message{}    
	
	msgChan, _ := pubsub.Sub()
	for msg := range msgChan {
			log.Info(" [x] %q", msg)
			json.Unmarshal([]byte(msg), res)
			policy, err1 := policies.GetPolicy("bind")
            if err1 != nil {
	           log.Error("Error: Policy :\n%s.", err1)
	         }
            
            asm, err := policies.GetAssembly(res.Id)
	        if err!= nil {
		       log.Error(err)
	        }
            
	        _, err2 := policy.Apply(asm)
	        if err2 != nil {
	            log.Error("Error: Policy doesn't apply :\n%s.", err2)
	         }
	         go app.RestartApp(asm)
			
		}
	log.Info("Handling message %v", msgChan)
	self.chann = msgChan
	
	//self.Serve()
}



