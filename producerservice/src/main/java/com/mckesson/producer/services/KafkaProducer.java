package com.mckesson.producer.services;

import com.mckesson.producer.entities.Message;
import com.mckesson.producer.utilities.Utilities;

import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
public class KafkaProducer {
    private static final Logger logger = LoggerFactory.getLogger(KafkaProducer.class);

    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    public void sendMessage(Message message) {
        logger.info(String.format("$$ -> Producing message --> %s", message));

        if (StringUtils.isNotBlank(message.getAppName()) && StringUtils.isNotBlank(message.getIncomingMessage())) {
            String KAFKA_TOPIC=null;

            String applicationName = message.getAppName().toUpperCase();

            switch (applicationName) {
                case "DRGPAYER":
                    KAFKA_TOPIC = Utilities.environmentOrDefault("DRG_PAYER_TOPIC_NAME", "default-topic");
                    break;
                case "DRGPLAN":
                    KAFKA_TOPIC = Utilities.environmentOrDefault("DRG_PLAN_TOPIC_NAME", "default-topic");
                    break;
                case "NCPDP":
                    KAFKA_TOPIC = Utilities.environmentOrDefault("NCPDP_TOPIC_NAME", "default-topic");
                    break;

            }

            if(StringUtils.isNotBlank(KAFKA_TOPIC)){
                this.kafkaTemplate.send(KAFKA_TOPIC, message.getIncomingMessage());
                logger.info("Message Recieved......   " + " Topic Name:::" + KAFKA_TOPIC + "Incoming Message: "
                        + message.getIncomingMessage());
            }else {
                logger.error("Failed to send the message....The Application is not supported...........");
            }
        }

    }
}