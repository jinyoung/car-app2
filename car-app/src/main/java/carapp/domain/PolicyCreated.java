package carapp.domain;

import carapp.domain.*;
import carapp.infra.AbstractEvent;
import java.util.*;
import lombok.*;


@Data
@ToString
public class PolicyCreated extends AbstractEvent {

    private Long id;
    private String policyId;
    private String carId;

    public PolicyCreated(PolicyApplication aggregate){
        super(aggregate);
    }
    public PolicyCreated(){
        super();
    }
}
