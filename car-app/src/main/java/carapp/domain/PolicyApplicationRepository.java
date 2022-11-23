package carapp.domain;

import carapp.domain.*;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource(collectionResourceRel="policyApplications", path="policyApplications")
public interface PolicyApplicationRepository extends PagingAndSortingRepository<PolicyApplication, Long>{

}
