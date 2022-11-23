package carapp.infra;
import carapp.domain.*;

import org.springframework.hateoas.Link;
import org.springframework.hateoas.server.RepresentationModelProcessor;
import org.springframework.stereotype.Component;
import org.springframework.hateoas.EntityModel;

@Component
public class PolicyApplicationHateoasProcessor implements RepresentationModelProcessor<EntityModel<PolicyApplication>>  {

    @Override
    public EntityModel<PolicyApplication> process(EntityModel<PolicyApplication> model) {

        
        return model;
    }
    
}
