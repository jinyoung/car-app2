package carapp.domain;

import carapp.domain.PolicyCreated;
import carapp.CarAppApplication;
import javax.persistence.*;
import java.util.List;
import lombok.Data;
import java.util.Date;


@Entity
@Table(name="PolicyApplication_table")
@Data

public class PolicyApplication  {


    
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    
    
    
    
    
    private Long id;
    
    
    
    
    
    private String policyId;
    
    
    
    
    
    private String carId;

    @PostPersist
    public void onPostPersist(){


        PolicyCreated policyCreated = new PolicyCreated(this);
        policyCreated.publishAfterCommit();

    }

    public static PolicyApplicationRepository repository(){
        PolicyApplicationRepository policyApplicationRepository = CarAppApplication.applicationContext.getBean(PolicyApplicationRepository.class);
        return policyApplicationRepository;
    }






}
