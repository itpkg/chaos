package com.itpkg.core.controllers;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

/**
 * Created by flamen on 16-5-28.
 */
//@EnableWebMvc
//@ControllerAdvice
public class ErrorController extends ResponseEntityExceptionHandler {
//    @ExceptionHandler(Exception.class)
//    void handleConflict(HttpServletResponse res, Exception ex) throws IOException{
//        logger.debug("################################");
//        res.getWriter().write("####################");
//        //return new ResponseEntity<>(ex.getMessage(), new HttpHeaders(), HttpStatus.INTERNAL_SERVER_ERROR);
//    }

//    @ExceptionHandler(Throwable.class)
//    @ResponseBody
//    ResponseEntity<Object> handleControllerException(HttpServletRequest req, Throwable ex) {
////        ErrorResponse errorResponse = new ErrorResponse(ex);
////        if(ex instanceof ServiceException) {
////            errorResponse.setDetails(((ServiceException)ex).getDetails());
////        }
////        if(ex instanceof ServiceHttpException) {
////            return new ResponseEntity<Object>(errorResponse,((ServiceHttpException)ex).getStatus());
////        } else {
////            return new ResponseEntity<Object>(errorResponse,HttpStatus.INTERNAL_SERVER_ERROR);
////        }
//        logger.debug("################# "+ex.getMessage());
//        return new ResponseEntity<Object>(ex.getMessage(),HttpStatus.INTERNAL_SERVER_ERROR);
//    }
//
//    @Override
//    protected ResponseEntity<Object> handleNoHandlerFoundException(NoHandlerFoundException ex, HttpHeaders headers, HttpStatus status, WebRequest request) {
//        Map<String,String> responseBody = new HashMap<>();
//        responseBody.put("path",request.getContextPath());
//        responseBody.put("message","The URL you have reached is not in service at this time (404).");
//        return new ResponseEntity<Object>(responseBody,HttpStatus.NOT_FOUND);
//    }

    private final static Logger logger = LoggerFactory.getLogger(ErrorController.class);
}
