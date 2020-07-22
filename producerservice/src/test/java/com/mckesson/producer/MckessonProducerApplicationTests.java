package com.mckesson.producer;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

import com.mckesson.producer.config.KafkaConfiguration;
import com.mckesson.producer.controller.MckessonController;
import com.mckesson.producer.entities.Message;

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

	@Autowired
	private KafkaConfiguration kafkaConfiguration;

	private static final Logger log = LoggerFactory.getLogger(MckessonProducerApplicationTests.class);

/* ================= Test Cases for  MckessonController====================*/
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

	@Test
	public void testProduce() {
		log.info("Test Rest end point URL produce .......................... .......................... ..........................");
		Message message = new Message();
		message.setTopicName("MyTopicName");
		message.setIncomingMessage( "My IncomingMessage");;
		mckController.produce(message);
		log.info("/kafka/produce end point with json  ....................................................: Success");

	}

	/* ================= MckessonController KafkaConfiguration ====================*/
	@Test
	public void testProducerFactory() {
		log.info("Test KafkaConfiguration producerFactory......." +kafkaConfiguration.producerFactory());
		assertNotNull(kafkaConfiguration.producerFactory());
	}

	@Test
	public void testProducerConfigurations() {
		log.info("Test KafkaConfiguration producerConfigurations. ................................" +kafkaConfiguration.producerConfigurations());
		assertNotNull(kafkaConfiguration.producerConfigurations());
	}

	@Test
	public void testKafkaTemplate() {
		log.info("Test KafkaConfiguration kafkaTemplate.... ............................." +kafkaConfiguration.kafkaTemplate());
		assertNotNull(kafkaConfiguration.kafkaTemplate());
	}

}
