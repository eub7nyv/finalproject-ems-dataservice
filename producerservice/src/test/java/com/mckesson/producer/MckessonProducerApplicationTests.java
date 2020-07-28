package com.mckesson.producer;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

import com.mckesson.producer.config.KafkaConfiguration;
import com.mckesson.producer.controller.MckessonController;
import com.mckesson.producer.entities.Message;
import com.mckesson.producer.services.KafkaProducer;

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

	@Autowired
	private KafkaProducer kafkaProducer;


	private static final Logger log = LoggerFactory.getLogger(MckessonProducerApplicationTests.class);

/* ================= Test Cases for  MckessonController====================*/
	@Test
	public void testHome() {
		log.info("1...............Test Default end point....." +mckController.home());
		assertNotNull(mckController.test());
	}

	@Test
	public void testTest() {
		log.info("2...............Test Rest end point URL......." +mckController.test());
		assertNotNull(mckController.test());
	}

	@Test
	public void testProduce() {
		Message message = new Message();
		message.setAppName("MyTopicName");
		message.setIncomingMessage( "My IncomingMessage");;
		mckController.produce(message);
		log.info("3.............../kafka/produce end point with json  ....................................................: Success");

	}

	/* ================= Test Cases for  KafkaConfiguration ====================*/
	@Test
	public void testProducerFactory() {
		log.info("4...............Test KafkaConfiguration producerFactory......." +kafkaConfiguration.producerFactory());
		assertNotNull(kafkaConfiguration.producerFactory());
	}

	@Test
	public void testProducerConfigurations() {
		log.info("5...............Test KafkaConfiguration producerConfigurations. ................................" +kafkaConfiguration.producerConfigurations());
		assertNotNull(kafkaConfiguration.producerConfigurations());
	}

	@Test
	public void testKafkaTemplate() {
		log.info("6...............Test KafkaConfiguration kafkaTemplate.... ............................." +kafkaConfiguration.kafkaTemplate());
		assertNotNull(kafkaConfiguration.kafkaTemplate());
	}

	/* ================= Test Cases for  KafkaProducer ====================*/
	//@Test
	public void testSendMessageWithValidApplicaionName() {
		Message message = new Message();
		message.setAppName("drgpayer"); //valid application name
		message.setIncomingMessage("My IncomingMessage from test case");
		kafkaProducer.sendMessage(message);
		log.info("7.............../mckesson/produce end point with json ..........................: Success: Good Application name");
	}

	@Test
	public void testSendMessageWithUnsupportedApplicaionName() {
		Message message = new Message();
		message.setAppName("Unsupported_application_name");//In-valid application name
		message.setIncomingMessage("My IncomingMessage from test case");
		kafkaProducer.sendMessage(message);
		log.info("8.............../mckesson/produce end point with json ..........................: Failed: Bad application name");
	}

}
