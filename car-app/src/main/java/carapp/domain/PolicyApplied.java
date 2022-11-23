package carapp.domain;

import carapp.domain.*;
import carapp.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class PolicyApplied extends AbstractEvent {

    private Long id;
    private Long carId;
    private Long policyApplicationId;
    private String status;
}
