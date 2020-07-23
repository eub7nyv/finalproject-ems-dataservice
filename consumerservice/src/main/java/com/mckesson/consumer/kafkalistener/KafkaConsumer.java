package com.mckesson.consumer.kafkalistener;

import com.mckesson.consumer.kafkaconfig.KafkaConfiguration;
import com.mckesson.consumer.service.ConsumerService;
import com.mckesson.consumer.utils.ConsumerConstants;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class KafkaConsumer {

    private final static Logger log = LoggerFactory.getLogger(KafkaConfiguration.class);

    @Autowired
    ConsumerService consumerService;


       @KafkaListener(topics="time-topic" , groupId = "docker-compose-consumer")
       public String  consumerExample(String message){
    	   System.out.println("<<<<<< message >>>>"+message);
        String statusMessage= consumerService.callDataServiceAPI(ConsumerConstants.TOPIC_TIME_TOPIC, message);
        System.out.println("<<<<<< Status Message From Service >>>>"+statusMessage);
        return statusMessage;
       }

        @KafkaListener(topics="drgpayer" , groupId = "docker-compose-consumer")
        public String consumerDRGPayerData(String message){
            log.info("Inside  consumerDRGPayerData >>>>"+message);
            System.out.println("<<<<<< message >>>>"+message);
            System.out.println("<<<<<< Status consumerDRGPayerData From Service >>>>"+message);
            String statusMessage= consumerService.callDataServiceAPI(ConsumerConstants.TOPIC_DRG_PAYER, message);
            System.out.println("<<<<<< Status Message From Service >>>>"+statusMessage);
            return statusMessage;
        }


        @KafkaListener(topics="drgplan" , groupId = "docker-compose-consumer")
        public String consumerDRGPlanData(String message){
            log.info("Inside  consumerDRGPayerData >>>>"+message);
            String statusMessage= consumerService.callDataServiceAPI(ConsumerConstants.TOPIC_DRG_PLAN, message);
            System.out.println("<<<<<< Status Message From Service >>>>"+statusMessage);
            return message;
        }


        @KafkaListener(topics="ncpdp" , groupId = "docker-compose-consumer")
        public String consumerNdpDpData(String message){
        	String statusMessage= consumerService.callDataServiceAPI(ConsumerConstants.TOPIC_NCPDP, message);
        	System.out.println("<<<<<< Status Message From Service >>>>"+statusMessage);
            return message;
        }



}
