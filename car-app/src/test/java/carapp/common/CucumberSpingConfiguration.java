package carapp.common;


import carapp.CarAppApplication;
import io.cucumber.spring.CucumberContextConfiguration;
import org.springframework.boot.test.context.SpringBootTest;

@CucumberContextConfiguration
@SpringBootTest(classes = { CarAppApplication.class })
public class CucumberSpingConfiguration {
    
}
