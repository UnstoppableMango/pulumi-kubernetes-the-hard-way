// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class EtcdConfigurationProps {
    /**
     * @return Path to the certificate authority file on the remote system.
     * 
     */
    private String caFilePath;
    /**
     * @return Path to the certificate file on the remote system.
     * 
     */
    private String certFilePath;
    /**
     * @return Etcd&#39;s data directory.
     * 
     */
    private String dataDirectory;
    /**
     * @return Path to the etcd binary.
     * 
     */
    private String etcdPath;
    /**
     * @return Internal IP of the etcd node.
     * 
     */
    private String internalIp;
    /**
     * @return Path to the private key file on the remote system.
     * 
     */
    private String keyFilePath;
    /**
     * @return Name of the etcd node.
     * 
     */
    private String name;

    private EtcdConfigurationProps() {}
    /**
     * @return Path to the certificate authority file on the remote system.
     * 
     */
    public String caFilePath() {
        return this.caFilePath;
    }
    /**
     * @return Path to the certificate file on the remote system.
     * 
     */
    public String certFilePath() {
        return this.certFilePath;
    }
    /**
     * @return Etcd&#39;s data directory.
     * 
     */
    public String dataDirectory() {
        return this.dataDirectory;
    }
    /**
     * @return Path to the etcd binary.
     * 
     */
    public String etcdPath() {
        return this.etcdPath;
    }
    /**
     * @return Internal IP of the etcd node.
     * 
     */
    public String internalIp() {
        return this.internalIp;
    }
    /**
     * @return Path to the private key file on the remote system.
     * 
     */
    public String keyFilePath() {
        return this.keyFilePath;
    }
    /**
     * @return Name of the etcd node.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EtcdConfigurationProps defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String caFilePath;
        private String certFilePath;
        private String dataDirectory;
        private String etcdPath;
        private String internalIp;
        private String keyFilePath;
        private String name;
        public Builder() {}
        public Builder(EtcdConfigurationProps defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.caFilePath = defaults.caFilePath;
    	      this.certFilePath = defaults.certFilePath;
    	      this.dataDirectory = defaults.dataDirectory;
    	      this.etcdPath = defaults.etcdPath;
    	      this.internalIp = defaults.internalIp;
    	      this.keyFilePath = defaults.keyFilePath;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder caFilePath(String caFilePath) {
            if (caFilePath == null) {
              throw new MissingRequiredPropertyException("EtcdConfigurationProps", "caFilePath");
            }
            this.caFilePath = caFilePath;
            return this;
        }
        @CustomType.Setter
        public Builder certFilePath(String certFilePath) {
            if (certFilePath == null) {
              throw new MissingRequiredPropertyException("EtcdConfigurationProps", "certFilePath");
            }
            this.certFilePath = certFilePath;
            return this;
        }
        @CustomType.Setter
        public Builder dataDirectory(String dataDirectory) {
            if (dataDirectory == null) {
              throw new MissingRequiredPropertyException("EtcdConfigurationProps", "dataDirectory");
            }
            this.dataDirectory = dataDirectory;
            return this;
        }
        @CustomType.Setter
        public Builder etcdPath(String etcdPath) {
            if (etcdPath == null) {
              throw new MissingRequiredPropertyException("EtcdConfigurationProps", "etcdPath");
            }
            this.etcdPath = etcdPath;
            return this;
        }
        @CustomType.Setter
        public Builder internalIp(String internalIp) {
            if (internalIp == null) {
              throw new MissingRequiredPropertyException("EtcdConfigurationProps", "internalIp");
            }
            this.internalIp = internalIp;
            return this;
        }
        @CustomType.Setter
        public Builder keyFilePath(String keyFilePath) {
            if (keyFilePath == null) {
              throw new MissingRequiredPropertyException("EtcdConfigurationProps", "keyFilePath");
            }
            this.keyFilePath = keyFilePath;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("EtcdConfigurationProps", "name");
            }
            this.name = name;
            return this;
        }
        public EtcdConfigurationProps build() {
            final var _resultValue = new EtcdConfigurationProps();
            _resultValue.caFilePath = caFilePath;
            _resultValue.certFilePath = certFilePath;
            _resultValue.dataDirectory = dataDirectory;
            _resultValue.etcdPath = etcdPath;
            _resultValue.internalIp = internalIp;
            _resultValue.keyFilePath = keyFilePath;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
