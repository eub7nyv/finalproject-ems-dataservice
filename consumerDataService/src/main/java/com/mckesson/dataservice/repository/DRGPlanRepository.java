package com.mckesson.dataservice.repository;

import com.mckesson.dataservice.bean.DRGPlanBean;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DRGPlanRepository extends CrudRepository<DRGPlanBean, Long> {

}
