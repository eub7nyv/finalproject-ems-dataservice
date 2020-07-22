package com.mckesson.producer.entities;

import org.springframework.boot.autoconfigure.domain.EntityScan;

@EntityScan
public class Message {

    private String topicName;
    private String incomingMessage;

    public String getTopicName() {
        return topicName;
    }

    public void setTopicName(String topicName) {
        this.topicName = topicName;
    }

    public String getIncomingMessage() {
        return incomingMessage;
    }

    public void setIncomingMessage(String incomingMessage) {
        this.incomingMessage = incomingMessage;
    }
}