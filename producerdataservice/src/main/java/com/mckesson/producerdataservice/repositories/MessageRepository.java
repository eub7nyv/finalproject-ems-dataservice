package com.mckesson.producerdataservice.repositories;

import com.mckesson.producerdataservive.entities.Message;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

/**
 * @author anoopunnikrishnan
 *
 */
@Repository
public interface MessageRepository extends CrudRepository<Message, Long> {

}