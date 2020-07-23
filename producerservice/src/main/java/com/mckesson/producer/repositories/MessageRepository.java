package com.mckesson.producer.repositories;

import com.mckesson.producer.entities.Payload;

import org.springframework.data.repository.CrudRepository;

public interface MessageRepository extends CrudRepository<Payload, Long> {

}