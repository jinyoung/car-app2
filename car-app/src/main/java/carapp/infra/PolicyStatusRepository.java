package carapp.infra;

import carapp.domain.*;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import java.util.List;

@RepositoryRestResource(collectionResourceRel="policyStatuses", path="policyStatuses")
public interface PolicyStatusRepository extends PagingAndSortingRepository<PolicyStatus, Long> {

    

    
}
