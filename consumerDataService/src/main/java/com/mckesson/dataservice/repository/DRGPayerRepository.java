package com.mckesson.dataservice.repository;

import com.mckesson.dataservice.bean.DRGPayerBean;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DRGPayerRepository extends CrudRepository<DRGPayerBean, Long> {
}

