package com.mckesson.producer.controller;

import com.mckesson.producer.entities.Message;
import com.mckesson.producer.services.KafkaProducer;
import com.mckesson.producer.utilities.Utilities;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

/**
 * @author anoopunnikrishnan
 *
 */

@Slf4j
@RestController
public class MckessonController {


    @Autowired
    private KafkaProducer kafkaProducer;

    @RequestMapping("/")
    public String home() {
        return "Hello Kafka Producer..";
    }

    @GetMapping("/test")
    public String test() {
        return "McKesson Producer controller is up and running.....";
    }

    @PostMapping(path = "/mckesson/produce")
    public void produce(@RequestBody Message message) {

        String streamingPlatform = Utilities.environmentOrDefault("STREAMING_PLATFORM", "KAFKA");
        log.info("Streaming message Platform :::::: {}",streamingPlatform);

        if (streamingPlatform.equals("KAFKA")) {
            kafkaProducer.sendMessage(message);
            log.info("Message Recieved With Application Name===> {} & Incoming Message====> {}" ,message.getAppName(), message.getIncomingMessage());
        } else if (streamingPlatform.equals("RABBIT")) {
            // TODO : Phase 2, Rabbit implimentation

        }

    }

}