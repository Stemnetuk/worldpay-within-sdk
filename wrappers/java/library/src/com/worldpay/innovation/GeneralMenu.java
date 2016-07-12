/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.PaymentResponse;
import com.worldpay.innovation.wpwithin.rpc.types.Price;
import com.worldpay.innovation.wpwithin.rpc.types.Service;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceDetails;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceMessage;
import com.worldpay.innovation.wpwithin.rpc.types.TotalPriceResponse;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.Scanner;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.thrift.TException;


/**
 *
 * @author worldpay
 */
public class GeneralMenu {
    private static final Logger log= Logger.getLogger( GeneralMenu.class.getName() );

    // TODO: put these somewhere sensible
    // TODO: What do these do and what should they be?
    private final String DEFAULT_DEVICE_NAME = "conorhwp-macbook";
    private final String DEFAULT_DEVICE_DESCRIPTION = "Conor H WP - Raspberry Pi";
    
    
    // Move these
    String ERR_DEVICE_NOT_INITIALISED = "ERR_DEVICE_NOT_INITIALISED";
    
    
    private final WPWithin.Client sdk;
    
    public GeneralMenu(WPWithin.Client _client) {
        log.setLevel(Level.FINE);
        ConsoleHandler handler = new ConsoleHandler();
        handler.setLevel(Level.FINE);
        log.addHandler(handler);
        this.sdk = _client;
    }
    
    public MenuReturnStruct mGetDeviceInfo() {

            if(this.sdk == null) {
                    return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
            }

            try {
                System.out.println("Uid of device: " + sdk.getDevice().getUid() + "\n");

                System.out.println("Name of device: " + sdk.getDevice().getName() + "\n");
                System.out.println("Description: " + sdk.getDevice().getDescription() + "\n");
                System.out.println("Services: \n");

                for(int i=0; i<sdk.getDevice().getServicesSize(); i++) {

                        Service service = (Service)((ArrayList)sdk.getDevice().getServices()).get(i);
                    
                        System.out.println("   " + i + ": Id:" + service.getId() + " Name:" + service.getName() + " Description:" + service.getDescription() + "\n");
                        System.out.println("   Prices: \n");
                        
                        
                        for(int j=0; j< service.getPricesSize(); j++) {
                            Price price = (Price)(((ArrayList)service.getPrices()).get(j));
                            System.out.println("      " + j + ": ServiceID: " + price.getServiceId() + " ID:" + price.getId() + " Description:" + price.getDescription() + " PricePerUnit:" + price.getPricePerUnit() + " UnitID:" + price.getUnitDescription() + " UnitDescription:%s\n");
                        }
                }

                System.out.println("IPv4Address: " + sdk.getDevice().getIpv4Address() + "\n");
                System.out.println("CurrencyCode: " + sdk.getDevice().getCurrencyCode() + "\n");

                return new MenuReturnStruct(null, 0);
                        
            } catch (TException ex) {
                Logger.getLogger(GeneralMenu.class.getName()).log(Level.SEVERE, "sdk client call failed", ex);
                return new MenuReturnStruct("sdk client call failed", 1);
            }                        
    }    

    public MenuReturnStruct mInitDefaultDevice() {

            //_sdk, err := wpwithin.Initialise(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION)
            
            try {
                sdk.setup(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION);
            } catch(TException e) {
                return new MenuReturnStruct("SDK setup failed", 1);
            }
        
            return new MenuReturnStruct(null, 0);

    }

    public MenuReturnStruct mInitNewDevice()  {

            System.out.println("Name of device: ");
            
            Scanner scanner = new Scanner(System.in);
            String nameOfDevice = scanner.next();
            if(null == nameOfDevice || "".equals(nameOfDevice)) {
                    return new MenuReturnStruct("Name of device not set", 0);
            }

            System.out.println("Description: ");
            String description = scanner.next();
            if(null == description || "".equals(description)) {
                    return new MenuReturnStruct("Description of device not set", 0);
            }
                    
            try {
                sdk.setup(nameOfDevice, description);
            } catch(TException e) {
                return new MenuReturnStruct("Setup of device unsucessful", 0);
            }
            
            return new MenuReturnStruct(null, 0);
            
    }

    
    public MenuReturnStruct mCarWashDemoConsumer() {

	log.fine("testDiscoveryAndNegotiation");

	MenuReturnStruct rc = mInitDefaultDevice();
        if(rc.getMsg() != null) {
            return rc;
        }    

        ConsumerMenu consumerMenu = new ConsumerMenu(sdk);
        
        rc = consumerMenu.mDefaultHCECredential();
        if(rc.getMsg() != null) {
            return rc;
        }    
        
	if(sdk == null) {
		return new MenuReturnStruct("ERR_DEVICE_NOT_INITIALISED", 0);
	}

        HashSet services;
	log.fine("pre scan for services");
	try {
            services = (HashSet)sdk.serviceDiscovery(20000);
        } catch(TException e) {
            return new MenuReturnStruct("Something failed during service discovery", 0);
        }
	log.fine("end scan for services");


	if(services.size() >= 1) {

		ServiceMessage svc = (ServiceMessage)(services.toArray()[0]);

		System.out.println("# Service:: (" + svc.getHostname() + ":" + svc.getPortNumber() + "/" + svc.getUrlPrefix() + ") - " + svc.getDeviceDescription());

		log.fine("Init consumer");
                        
                try {        
                    sdk.initConsumer("http://", svc.getHostname(), svc.getPortNumber(), svc.getUrlPrefix(), svc.getServerId());
                } catch(TException e) {
                    return new MenuReturnStruct("Faild to init the consumer", 0);
                }

		log.fine("Client created..");
                ArrayList serviceDetails;
                
                try {
                    serviceDetails = (ArrayList)sdk.requestServices();
                } catch(TException e) {
                    return new MenuReturnStruct("Failed to request services", 0);
                }

		if(serviceDetails.size() >= 1) {

			ServiceDetails svcDetails = (ServiceDetails)serviceDetails.get(0);

			System.out.println(svcDetails.getServiceId() + " - " + svcDetails.getServiceDescription() + "\n");

                        ArrayList prices;
                        try {
        			prices = (ArrayList)sdk.getServicePrices(svcDetails.getServiceId());
                            
                        } catch(TException e) {
                            return new MenuReturnStruct("Failed to get prices", 0);
                        }

			System.out.println("------- Prices -------\n");
			if(prices.size() >= 1) {

				Price price = (Price)prices.get(0);

				System.out.println("(" + price.getId() + ") " + price.getDescription() + " @ " + price.getPricePerUnit() + ", " + price.getUnitDescription() + " (Unit id = " + price.getUnitId() +")\n");

                                TotalPriceResponse tpr;
                                try {

                                    tpr = sdk.selectService(price.getServiceId(), 2, price.getId());
                                    
                                } catch(TException e) {
                                    return new MenuReturnStruct("Failed to get total price response", 0);
                                }
                                
                                System.out.println("#Begin Request#");
				System.out.println("ServerID: " + tpr.getServerId() + "\n");
				System.out.println("PriceID = " + tpr.getPriceId() + " - " + tpr.getUnitsToSupply() + " units = " + tpr.getTotalPrice() + "\n");
				System.out.println("ClientID: " + tpr.getClientId() + ", MerchantClientKey: " + tpr.getMerchantClientKey() + ", PaymentRef: " + tpr.getPaymentReferenceId() + "\n");
				System.out.println("#End Request#");

				log.log(Level.FINE, "Making payment of {0}\n", tpr.getTotalPrice());

                                PaymentResponse payResp;
                                try {
                                    payResp = sdk.makePayment(tpr);
                                } catch(TException e) {
                                    return new MenuReturnStruct("Failed to make the payment unfortunately", 0);
                                }
                                
				System.out.println("Payment of " + payResp.getTotalPaid() + " made successfully\n");

				System.out.println("Service delivery token: " + payResp.getServiceDeliveryToken() + "\n");

			}
		}
	}
	return new MenuReturnStruct(null, 0);
}

    
public MenuReturnStruct mCarWashDemoProducer() {

    ProducerMenu producerMenu = new ProducerMenu(this.sdk);
    return producerMenu.mCarWashDemoProducer();
    
}
            
            
    /*
func mResetSessionState() (int, error) {

	sdk = nil

	return 0, nil
}
*/
    
    /*
func mLoadConfig() (int, error) {

	// Ask user for path to config file
	// (And password if secured)

	return 0, errors.New("Not implemented yet..")
}
*/
    
    /*
func mReadConfig() (int, error) {

	// Print out loaded configuration
	// Print out the path to file that was loaded (Need to keep reference during load stage)

	return 0, errors.New("Not implemented yet..")
}
*/
    /*
func mStartRPCService() (int, error) {

	config := rpc.Configuration{
		Protocol:   "binary",
		Framed:     false,
		Buffered:   false,
		Host:       "127.0.0.1",
		Port:       9091,
		Secure:     false,
		BufferSize: 8192,
	}

	rpc, err := rpc.NewService(config, sdk)

	if err != nil {

		return 0, err
	}

	if err := rpc.Start(); err != nil {

		return 0, err
	}

	return 0, nil
}
*/


}
