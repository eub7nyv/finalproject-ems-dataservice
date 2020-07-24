package com.mckesson.producer.services;

import java.util.Optional;

import com.mckesson.producer.dto.PayloadDTO;
import com.mckesson.producer.entities.Payload;
import com.mckesson.producer.repositories.MessageRepository;
import com.mckesson.producer.utilities.ObjectMapperUtils;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author anoopunnikrishnan
 *
 */

 
@Service
public class MessageService {

    @Autowired
    private MessageRepository messageRepository;
    

    public PayloadDTO save(PayloadDTO payloadDTO) {
		Optional<Payload> incomingMessage = messageRepository.findById(payloadDTO.getId());
		Payload payload = incomingMessage.get();
		if (payload != null) {
			payload.setApplicationName(payloadDTO.getApplicationName());
			payload.setIncomingMessage(payloadDTO.getIncomingMessage());


			payload = messageRepository.save(payload);
		}
		return (payload == null ? null : ObjectMapperUtils.map(payload, PayloadDTO.class));
	}
    
}