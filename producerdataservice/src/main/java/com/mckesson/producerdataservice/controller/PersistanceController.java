package com.mckesson.producerdataservice.controller;

import com.mckesson.producerdataservice.service.MessageService;


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PersistanceController {

    private static final Logger log = LoggerFactory.getLogger(PersistanceController.class);

    @Autowired
    MessageService messageService;

	@GetMapping("/test")
    public String greeting() {
        return "Hello From  PersistanceController";
    }
    
}