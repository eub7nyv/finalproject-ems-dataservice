package com.mckesson.dataservice.service;

import com.mckesson.dataservice.bean.DRGPayerBean;
import com.mckesson.dataservice.bean.DRGPlanBean;
import com.mckesson.dataservice.bean.NCPDPBean;
import com.mckesson.dataservice.controller.DataController;
import com.mckesson.dataservice.entity.ConsumerConstants;
import com.mckesson.dataservice.repository.DRGPayerRepository;
import com.mckesson.dataservice.repository.DRGPlanRepository;
import com.mckesson.dataservice.repository.NCPDPRepository;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.StringTokenizer;
import java.util.stream.Collectors;

@Service
public class ConsumerDataService {

	private static final Logger log = LoggerFactory.getLogger(ConsumerDataService.class);

	@Autowired
	DRGPayerRepository consumerPayerRepository;

	@Autowired
	DRGPlanRepository consumerPlanRepository;
	
	@Autowired
	NCPDPRepository ncpdpPRepository;
	

	public String insertDRGPayerConsumerData(String message) {
		log.info("In Side insertDRGPayerConsumerData >>>" + message);

		String[] arrayData = processMessageData(message);
		String statusMessage = ConsumerConstants.ERROR_MESSAGE;
		try {
			DRGPayerBean consumerPayerBean = new DRGPayerBean();
			if (arrayData != null && arrayData.length > 0) {
				consumerPayerBean.setPayer_Id(arrayData[0]);
				consumerPayerBean.setPayer_Name(arrayData[1]);
				consumerPayerBean.setParent_Id(arrayData[2]);
				consumerPayerBean.setCreatedBy(ConsumerConstants.TOPIC_DRG_PAYER);
				consumerPayerBean.setLastUpdateddBy(ConsumerConstants.TOPIC_DRG_PAYER);
			}
			consumerPayerBean = consumerPayerRepository.save(consumerPayerBean);
			statusMessage = ConsumerConstants.SUCCESS_MESSAGE;
		} catch (Exception e) {
			statusMessage=statusMessage+e.getMessage();
			log.error("Exception in saving data insertDRGPayerConsumerData"+e.getMessage());
		}

		return statusMessage;
	}

	public String insertDRGPlanConsumerData(String message) {

		log.info("In Side insertDRGPlanConsumerData >>>" + message);

		String[] arrayData = processMessageData(message);
		String statusMessage = ConsumerConstants.ERROR_MESSAGE;
		try {
			
			//PAYER_ID|Plan ID|Benefit Offering Name|SEGMENT|PLAN_TYPE_NAME|PPO|HMO|POS    
			DRGPlanBean consumerPlanBean = new DRGPlanBean();
			if (arrayData != null && arrayData.length > 0) {
				consumerPlanBean.setPayer_Id(arrayData[0]);
				consumerPlanBean.setPlan_Id(arrayData[1]);
				consumerPlanBean.setPlan_name(arrayData[2]);
				consumerPlanBean.setSegment(arrayData[3]);
				consumerPlanBean.setPlan_Type_Name(arrayData[4]);
				consumerPlanBean.setPpo(arrayData[5]);
				consumerPlanBean.setHmo(arrayData[6]);
				consumerPlanBean.setPos(arrayData[7]);
				consumerPlanBean.setCreatedBy(ConsumerConstants.TOPIC_DRG_PLAN);
				consumerPlanBean.setLastUpdateddBy(ConsumerConstants.TOPIC_DRG_PLAN);
			}
			consumerPlanBean = consumerPlanRepository.save(consumerPlanBean);
			statusMessage = ConsumerConstants.SUCCESS_MESSAGE;
		} catch (Exception e) {
			log.error("Exception in saving data insertDRGPlanConsumerData");
		}

		return statusMessage;
	}

	private String[]  processMessageData(String pipeDelimited) {
		
		return pipeDelimited.split("\\|");
	}
	
	

	public List<DRGPayerBean> getAllDRGPayerConsumerDataList() throws Exception {
		log.info("In Side getAllDRGPayerConsumerDataList >>>");
		final List<DRGPayerBean> drgpayerBeanList = new ArrayList<DRGPayerBean>();

		for (DRGPayerBean consumerBean : consumerPayerRepository.findAll()) {

			drgpayerBeanList.add(consumerBean);
		}
		log.info("In Side getAllDRGPayerConsumerDataList >>>" + drgpayerBeanList.size());
		return drgpayerBeanList;
	}

	public List<DRGPlanBean> getAllDRGPlanConsumerDataList() {
		log.info("In Side getAllDRGPlanConsumerDataList >>>");
		final List<DRGPlanBean> drgplanBeanList = new ArrayList<DRGPlanBean>();

		for (DRGPlanBean drgPlanBean : consumerPlanRepository.findAll()) {

			drgplanBeanList.add(drgPlanBean);
		}
		log.info("In Side getAllDRGPayerConsumerDataList >>>" + drgplanBeanList.size());
		return drgplanBeanList;
	}
	
	
	public List<NCPDPBean> getAllNCPDPConsumerDataList() {
		log.info("In Side getAllNCPDPConsumerDataList >>>");
		final List<NCPDPBean> ncpdpnBeanList = new ArrayList<NCPDPBean>();

		for (NCPDPBean ncpdpbean : ncpdpPRepository.findAll()) {

			ncpdpnBeanList.add(ncpdpbean);
		}
		log.info("In Side getAllDRGPayerConsumerDataList >>>" + ncpdpnBeanList.size());
		return ncpdpnBeanList;
	}

}
