package com.mckesson.producer.controller;

import org.springframework.web.bind.annotation.*;

@RestController
public class MckessonController {


    @RequestMapping("/")
	public String home() {
	  return "Hello Kafka Producer..";
    }

    @GetMapping("/test")
    public String test() {
        return "McKesson Producer controller is up and running.....";
    }

}