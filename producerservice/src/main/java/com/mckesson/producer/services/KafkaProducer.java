package com.mckesson.producer.services;

import com.mckesson.producer.entities.Message;
import com.mckesson.producer.utilities.Utilities;

import org.apache.commons.lang3.StringUtils;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;
import org.springframework.util.concurrent.ListenableFutureCallback;
import org.springframework.kafka.support.SendResult;

import org.springframework.lang.NonNull;

/**
 * @author anoopunnikrishnan
 *
 */

@Slf4j
@Service
public class KafkaProducer {


    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    public void sendMessage(Message message) {
        log.info(String.format("$$ -> Producing message --> {}", message));

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
            if(StringUtils.isNotBlank(KAFKA_TOPIC) && StringUtils.isNotBlank(message.getIncomingMessage())){

                final String kafkatopic=KAFKA_TOPIC;
                final String msg=message.getIncomingMessage();

                this.kafkaTemplate.send(KAFKA_TOPIC, message.getIncomingMessage())
                .addCallback(
                    new ListenableFutureCallback<SendResult<String, String>>() {
                      @Override
                      public void onFailure(@NonNull Throwable throwable) {
                        log.warn("Failed to send the message to Topic ===> {}  & the  Message send is ===> {} ",  kafkatopic,  msg, throwable);
                      }
          
                      @Override
                      public void onSuccess(SendResult<String, String> objectObjectSendResult) {
                        log.info("Message Send Successfully to Topic ===> {}  & the  Message send is ===> {} ", kafkatopic,  msg);
                      }
                    });
            }else {
                log.error("Failed to send the message....<Requested Application is not supported by this messaging service>");
            }
        }

    }
}
