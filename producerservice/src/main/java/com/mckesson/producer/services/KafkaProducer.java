package com.mckesson.producer.services;

import com.mckesson.producer.entities.Message;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
public class KafkaProducer {
    private static final Logger logger = LoggerFactory.getLogger(KafkaProducer.class);

    @Autowired
    private KafkaTemplate<String,String> kafkaTemplate;
    
    public void sendMessage(String topicName, Message message) {
        logger.info(String.format("$$ -> Producing message --> %s", message));
        this.kafkaTemplate.send(topicName, message.getIncomingMessage());
        logger.info("Message Recieved......   " + " Topic Name:::" + topicName + "Incoming Message: " + message.getIncomingMessage());
 
    }
}