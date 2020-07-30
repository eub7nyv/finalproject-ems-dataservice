package com.mckesson.producerdataservice.service;

import com.mckesson.producerdataservice.dto.MessageDTO;
import com.mckesson.producerdataservive.entities.Message;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * @author anoopunnikrishnan
 *
 */

@Service
public class MessageService {



    private static final Logger logger = LoggerFactory.getLogger(MessageService.class);

    public String save(Message message) {

        if (message != null) {
            //message = messageRepository.save(new Message(message.getApplicationName(), message.getIncomingMessage(), false));
            logger.info("Successfully Saved=============>" + message);
            return("Success");
        }else{
            logger.info("Save failed=============>" + message);
            return("Failed");
        }

    }

}