package com.mckesson.dataservice.exception;

import org.springframework.http.HttpStatus;
import org.springframework.validation.BindingResult;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;

import com.mckesson.dataservice.bean.DRGPayerBean;

@ControllerAdvice
public class ApiExceptionHandler {
 
    @ExceptionHandler(MethodArgumentNotValidException.class)
    @ResponseStatus(code = HttpStatus.BAD_REQUEST)
    @ResponseBody
    public String handleValidationError(MethodArgumentNotValidException ex) {
        BindingResult bindingResult = ex.getBindingResult();
        FieldError fieldError = bindingResult.getFieldError();
        String defaultMessage = fieldError.getDefaultMessage();
        return "BAD_REQUEST; Please check again";
    }
}