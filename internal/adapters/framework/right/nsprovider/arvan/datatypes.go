package arvan

import "time"

//////////////

type dominsResp struct {
	Data []struct {
		ID        string `json:"id"`
		UserID    string `json:"user_id"`
		Domain    string `json:"domain"`
		Name      string `json:"name"`
		PlanLevel int    `json:"plan_level"`
		Services  struct {
			Cdn           bool   `json:"cdn"`
			DNS           string `json:"dns"`
			CloudSecurity bool   `json:"cloud_security"`
		} `json:"services"`
		NsKeys       []string `json:"ns_keys"`
		CurrentNs    []string `json:"current_ns"`
		TargetCname  string   `json:"target_cname"`
		CustomCname  string   `json:"custom_cname"`
		Type         string   `json:"type"`
		Status       string   `json:"status"`
		DNSCloud     bool     `json:"dns_cloud"`
		IsPaused     bool     `json:"is_paused"`
		IsSuspended  bool     `json:"is_suspended"`
		ParentDomain bool     `json:"parent_domain"`
		Transfer     struct {
			Domain      string    `json:"domain"`
			AccountID   string    `json:"account_id"`
			AccountName string    `json:"account_name"`
			Time        time.Time `json:"time"`
		} `json:"transfer"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

type aRecResp struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value []struct {
			IP             string `json:"ip"`
			Port             interface{}    `json:"port"`
			Weight         int    `json:"weight"`
			OriginalWeight int    `json:"original_weight"`
			Country        string `json:"country"`
		} `json:"value"`
		TTL           int    `json:"ttl"`
		Cloud         bool   `json:"cloud"`
		UpstreamHTTPS string `json:"upstream_https"`
		IPFilterMode  struct {
			Count     string `json:"count"`
			Order     string `json:"order"`
			GeoFilter string `json:"geo_filter"`
		} `json:"ip_filter_mode"`
		CanDelete        bool      `json:"can_delete"`
		IsProtected      bool      `json:"is_protected"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		MonitoringStatus string    `json:"monitoring_status"`
		HealthCheck      struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			Description   string   `json:"description"`
			Origin        string   `json:"origin"`
			OriginType    string   `json:"origin_type"`
			Upstreams     []string `json:"upstreams"`
			Interval      int      `json:"interval"`
			Threshold     int      `json:"threshold"`
			Type          string   `json:"type"`
			Status        bool     `json:"status"`
			Retries       int      `json:"retries"`
			RequestConfig struct {
				Method           string `json:"method"`
				Port             interface{}    `json:"port"`
				Path             string `json:"path"`
				AllowInsecure    bool   `json:"allow_insecure"`
				ExpectedResponse struct {
					Codes   []int `json:"codes"`
					Headers struct {
						Property1 []string `json:"property1"`
						Property2 []string `json:"property2"`
					} `json:"headers"`
					Body string `json:"body"`
				} `json:"expected_response"`
				Headers struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"headers"`
				Timeout int `json:"timeout"`
			} `json:"request_config"`
			Zones []struct {
				ID              string `json:"id"`
				MonitoringLevel string `json:"monitoring_level"`
			} `json:"zones"`
			MonitoringUpdatedAt time.Time `json:"monitoring_updated_at"`
		} `json:"health_check"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

type cnameRecResp struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value struct {
			Host       string `json:"host"`
			HostHeader string `json:"host_header"`
			Port       interface{}   `json:"port"`
		} `json:"value"`
		TTL           int    `json:"ttl"`
		Cloud         bool   `json:"cloud"`
		UpstreamHTTPS string `json:"upstream_https"`
		IPFilterMode  struct {
			Count     string `json:"count"`
			Order     string `json:"order"`
			GeoFilter string `json:"geo_filter"`
		} `json:"ip_filter_mode"`
		CanDelete        bool      `json:"can_delete"`
		IsProtected      bool      `json:"is_protected"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		MonitoringStatus string    `json:"monitoring_status"`
		HealthCheck      struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			Description   string   `json:"description"`
			Origin        string   `json:"origin"`
			OriginType    string   `json:"origin_type"`
			Upstreams     []string `json:"upstreams"`
			Interval      int      `json:"interval"`
			Threshold     int      `json:"threshold"`
			Type          string   `json:"type"`
			Status        bool     `json:"status"`
			Retries       int      `json:"retries"`
			RequestConfig struct {
				Method           string `json:"method"`
				Port             interface{}    `json:"port"`
				Path             string `json:"path"`
				AllowInsecure    bool   `json:"allow_insecure"`
				ExpectedResponse struct {
					Codes   []int `json:"codes"`
					Headers struct {
						Property1 []string `json:"property1"`
						Property2 []string `json:"property2"`
					} `json:"headers"`
					Body string `json:"body"`
				} `json:"expected_response"`
				Headers struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"headers"`
				Timeout int `json:"timeout"`
			} `json:"request_config"`
			Zones []struct {
				ID              string `json:"id"`
				MonitoringLevel string `json:"monitoring_level"`
			} `json:"zones"`
			MonitoringUpdatedAt time.Time `json:"monitoring_updated_at"`
		} `json:"health_check"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

type txtRecResp struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value struct {
			Text string `json:"text"`
		} `json:"value"`
		TTL           int    `json:"ttl"`
		Cloud         bool   `json:"cloud"`
		UpstreamHTTPS string `json:"upstream_https"`
		IPFilterMode  struct {
			Count     string `json:"count"`
			Order     string `json:"order"`
			GeoFilter string `json:"geo_filter"`
		} `json:"ip_filter_mode"`
		CanDelete        bool      `json:"can_delete"`
		IsProtected      bool      `json:"is_protected"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		MonitoringStatus string    `json:"monitoring_status"`
		HealthCheck      struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			Description   string   `json:"description"`
			Origin        string   `json:"origin"`
			OriginType    string   `json:"origin_type"`
			Upstreams     []string `json:"upstreams"`
			Interval      int      `json:"interval"`
			Threshold     int      `json:"threshold"`
			Type          string   `json:"type"`
			Status        bool     `json:"status"`
			Retries       int      `json:"retries"`
			RequestConfig struct {
				Method           string `json:"method"`
				Port             interface{}    `json:"port"`
				Path             string `json:"path"`
				AllowInsecure    bool   `json:"allow_insecure"`
				ExpectedResponse struct {
					Codes   []int `json:"codes"`
					Headers struct {
						Property1 []string `json:"property1"`
						Property2 []string `json:"property2"`
					} `json:"headers"`
					Body string `json:"body"`
				} `json:"expected_response"`
				Headers struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"headers"`
				Timeout int `json:"timeout"`
			} `json:"request_config"`
			Zones []struct {
				ID              string `json:"id"`
				MonitoringLevel string `json:"monitoring_level"`
			} `json:"zones"`
			MonitoringUpdatedAt time.Time `json:"monitoring_updated_at"`
		} `json:"health_check"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

type mxRecResp struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value struct {
			Host     string `json:"host"`
			Priority interface{}    `json:"priority"`
		} `json:"value"`
		TTL           int    `json:"ttl"`
		Cloud         bool   `json:"cloud"`
		UpstreamHTTPS string `json:"upstream_https"`
		IPFilterMode  struct {
			Count     string `json:"count"`
			Order     string `json:"order"`
			GeoFilter string `json:"geo_filter"`
		} `json:"ip_filter_mode"`
		CanDelete        bool      `json:"can_delete"`
		IsProtected      bool      `json:"is_protected"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		MonitoringStatus string    `json:"monitoring_status"`
		HealthCheck      struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			Description   string   `json:"description"`
			Origin        string   `json:"origin"`
			OriginType    string   `json:"origin_type"`
			Upstreams     []string `json:"upstreams"`
			Interval      int      `json:"interval"`
			Threshold     int      `json:"threshold"`
			Type          string   `json:"type"`
			Status        bool     `json:"status"`
			Retries       int      `json:"retries"`
			RequestConfig struct {
				Method           string `json:"method"`
				Port             interface{}    `json:"port"`
				Path             string `json:"path"`
				AllowInsecure    bool   `json:"allow_insecure"`
				ExpectedResponse struct {
					Codes   []int `json:"codes"`
					Headers struct {
						Property1 []string `json:"property1"`
						Property2 []string `json:"property2"`
					} `json:"headers"`
					Body string `json:"body"`
				} `json:"expected_response"`
				Headers struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"headers"`
				Timeout int `json:"timeout"`
			} `json:"request_config"`
			Zones []struct {
				ID              string `json:"id"`
				MonitoringLevel string `json:"monitoring_level"`
			} `json:"zones"`
			MonitoringUpdatedAt time.Time `json:"monitoring_updated_at"`
		} `json:"health_check"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

type ptrRecResp struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value struct {
			Domain string `json:"domain"`
		} `json:"value"`
		TTL           int    `json:"ttl"`
		Cloud         bool   `json:"cloud"`
		UpstreamHTTPS string `json:"upstream_https"`
		IPFilterMode  struct {
			Count     string `json:"count"`
			Order     string `json:"order"`
			GeoFilter string `json:"geo_filter"`
		} `json:"ip_filter_mode"`
		CanDelete        bool      `json:"can_delete"`
		IsProtected      bool      `json:"is_protected"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		MonitoringStatus string    `json:"monitoring_status"`
		HealthCheck      struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			Description   string   `json:"description"`
			Origin        string   `json:"origin"`
			OriginType    string   `json:"origin_type"`
			Upstreams     []string `json:"upstreams"`
			Interval      int      `json:"interval"`
			Threshold     int      `json:"threshold"`
			Type          string   `json:"type"`
			Status        bool     `json:"status"`
			Retries       int      `json:"retries"`
			RequestConfig struct {
				Method           string `json:"method"`
				Port             interface{}   `json:"port"`
				Path             string `json:"path"`
				AllowInsecure    bool   `json:"allow_insecure"`
				ExpectedResponse struct {
					Codes   []int `json:"codes"`
					Headers struct {
						Property1 []string `json:"property1"`
						Property2 []string `json:"property2"`
					} `json:"headers"`
					Body string `json:"body"`
				} `json:"expected_response"`
				Headers struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"headers"`
				Timeout int `json:"timeout"`
			} `json:"request_config"`
			Zones []struct {
				ID              string `json:"id"`
				MonitoringLevel string `json:"monitoring_level"`
			} `json:"zones"`
			MonitoringUpdatedAt time.Time `json:"monitoring_updated_at"`
		} `json:"health_check"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

type nsRecResp struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value struct {
			Host string `json:"host"`
		} `json:"value"`
		TTL           int    `json:"ttl"`
		Cloud         bool   `json:"cloud"`
		UpstreamHTTPS string `json:"upstream_https"`
		IPFilterMode  struct {
			Count     string `json:"count"`
			Order     string `json:"order"`
			GeoFilter string `json:"geo_filter"`
		} `json:"ip_filter_mode"`
		CanDelete        bool      `json:"can_delete"`
		IsProtected      bool      `json:"is_protected"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		MonitoringStatus string    `json:"monitoring_status"`
		HealthCheck      struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			Description   string   `json:"description"`
			Origin        string   `json:"origin"`
			OriginType    string   `json:"origin_type"`
			Upstreams     []string `json:"upstreams"`
			Interval      int      `json:"interval"`
			Threshold     int      `json:"threshold"`
			Type          string   `json:"type"`
			Status        bool     `json:"status"`
			Retries       int      `json:"retries"`
			RequestConfig struct {
				Method           string `json:"method"`
				Port             interface{}    `json:"port"`
				Path             string `json:"path"`
				AllowInsecure    bool   `json:"allow_insecure"`
				ExpectedResponse struct {
					Codes   []int `json:"codes"`
					Headers struct {
						Property1 []string `json:"property1"`
						Property2 []string `json:"property2"`
					} `json:"headers"`
					Body string `json:"body"`
				} `json:"expected_response"`
				Headers struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"headers"`
				Timeout int `json:"timeout"`
			} `json:"request_config"`
			Zones []struct {
				ID              string `json:"id"`
				MonitoringLevel string `json:"monitoring_level"`
			} `json:"zones"`
			MonitoringUpdatedAt time.Time `json:"monitoring_updated_at"`
		} `json:"health_check"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

