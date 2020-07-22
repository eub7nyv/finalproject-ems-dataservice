package com.mckesson.producer.controller;

import com.mckesson.producer.entities.Message;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.*;

@RestController
public class MckessonController {

    private static final Logger log = LoggerFactory.getLogger(MckessonController.class);
    
    @RequestMapping("/")
	public String home() {
	  return "Hello Kafka Producer..";
    }

    @GetMapping("/test")
    public String test() {
        return "McKesson Producer controller is up and running.....";
    }

    @PostMapping(path = "/kafka/produce")
    public void produce(@RequestBody Message msg) {

        /* Wite the Kafka producer call here.
         * For time being I am printing the customer Message just to show the POST call is working.
         */

        log.info( "Message Recieved:::::   "+"Topic Name:::" +msg.getTopicName()+ "Incoming Message: " +msg.getIncomingMessage() ) ;
    }

}