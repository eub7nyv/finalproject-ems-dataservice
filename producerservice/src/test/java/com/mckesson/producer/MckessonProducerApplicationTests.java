package com.mckesson.producer;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

import com.mckesson.producer.controller.MckessonController;

import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.ComponentScan;

@ComponentScan("com.mckesson.producer")
@SpringBootTest
class MckessonProducerApplicationTests {

	@Test
	void contextLoads() {
	}

	@Autowired
	private MckessonController mckController;

	private static final Logger log = LoggerFactory.getLogger(MckessonProducerApplicationTests.class);

	@Test
	public void testHome() {
		log.info("Test Default end point....." +mckController.home());
		assertNotNull(mckController.test());
	}

	@Test
	public void testTest() {
		log.info("Test Rest end point URL......." +mckController.test());
		assertNotNull(mckController.test());
	}

}
