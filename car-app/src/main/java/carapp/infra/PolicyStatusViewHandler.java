package carapp.infra;

import carapp.domain.*;
import carapp.config.kafka.KafkaProcessor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.stream.annotation.StreamListener;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.List;
import java.util.Optional;

@Service
public class PolicyStatusViewHandler {

    @Autowired
    private PolicyStatusRepository policyStatusRepository;

    @StreamListener(KafkaProcessor.INPUT)
    public void whenPolicyCreated_then_CREATE_1 (@Payload PolicyCreated policyCreated) {
        try {

            if (!policyCreated.validate()) return;

            // view 객체 생성
            PolicyStatus policyStatus = new PolicyStatus();
            // view 객체에 이벤트의 Value 를 set 함
            policyStatus.setId(policyCreated.getId());
            policyStatus.setStatus("Created");
            // view 레파지 토리에 save
            policyStatusRepository.save(policyStatus);

        }catch (Exception e){
            e.printStackTrace();
        }
    }


    @StreamListener(KafkaProcessor.INPUT)
    public void whenPolicyApplied_then_UPDATE_1(@Payload PolicyApplied policyApplied) {
        try {
            if (!policyApplied.validate()) return;
                // view 객체 조회
            Optional<PolicyStatus> policyStatusOptional = policyStatusRepository.findById(policyApplied.getPolicyApplicationId());

            if( policyStatusOptional.isPresent()) {
                 PolicyStatus policyStatus = policyStatusOptional.get();
            // view 객체에 이벤트의 eventDirectValue 를 set 함
                policyStatus.setStatus("Applied");    
                // view 레파지 토리에 save
                 policyStatusRepository.save(policyStatus);
                }


        }catch (Exception e){
            e.printStackTrace();
        }
    }
    @StreamListener(KafkaProcessor.INPUT)
    public void whenPolicyDenied_then_UPDATE_2(@Payload PolicyDenied policyDenied) {
        try {
            if (!policyDenied.validate()) return;
                // view 객체 조회
            Optional<PolicyStatus> policyStatusOptional = policyStatusRepository.findById(policyDenied.getPolicyApplicationId());

            if( policyStatusOptional.isPresent()) {
                 PolicyStatus policyStatus = policyStatusOptional.get();
            // view 객체에 이벤트의 eventDirectValue 를 set 함
                policyStatus.setStatus("Denied");    
                // view 레파지 토리에 save
                 policyStatusRepository.save(policyStatus);
                }


        }catch (Exception e){
            e.printStackTrace();
        }
    }


}

