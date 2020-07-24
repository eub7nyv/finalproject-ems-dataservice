package com.mckesson.producer.repositories;

import com.mckesson.producer.entities.Payload;

import org.springframework.data.repository.CrudRepository;

/**
 * @author anoopunnikrishnan
 *
 */

public interface MessageRepository extends CrudRepository<Payload, Long> {

}