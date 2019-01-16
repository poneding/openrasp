package es

var attackAlarmTemplate = `
		{
			"template":"openrasp-attack-alarm-*",
			"aliases" : {
        		"real-{index}" : {} 
    		},
			"settings": {
				"analysis": {
					"normalizer": {
						"lowercase_normalizer": {
							"type": "custom",
							"filter": ["lowercase"]
						}
					}     
				}
			},
			"mappings": {
				"attack-alarm": {
					"_all": {
						"enabled": false
					},
					"properties": {
						"@timestamp":{
							"type":"date"
						},
						"request_method": {
							"type": "keyword",
							"ignore_above": 50
						},
						"target": {
							"type": "keyword",
							"ignore_above": 256
						},
						"server_ip": {
							"type": "keyword",
							"ignore_above": 256
						},
						"client_ip": {
							"type": "keyword",
							"ignore_above": 256
						},
						"referer": {
							"type": "keyword",
							"ignore_above": 256
						},
						"user_agent": {
							"type": "keyword",
							"ignore_above": 512
						},
						"attack_source": {
							"type": "keyword",
							"ignore_above": 256
						},
						"path": {
							"type": "keyword",
							"ignore_above": 256
						},
						"url": {
							"type": "keyword",
							"ignore_above": 256,
							"normalizer": "lowercase_normalizer"
						},
						"event_type": {
							"type": "keyword",
							"ignore_above": 256
						},
						"server_hostname": {
							"type": "keyword",
							"ignore_above": 256,
							"normalizer": "lowercase_normalizer"
						},
						"stack_md5": {
							"type": "keyword",
							"ignore_above": 64
						},
						"server_type": {
							"type": "keyword",
							"ignore_above": 256
						},
						"server_version": {
							"type": "keyword",
							"ignore_above": 256
						},
						"request_id": {
							"type": "keyword",
							"ignore_above": 256
						},
						"body": {
							"type": "keyword"
						},
						"app_id": {
							"type": "keyword",
							"ignore_above": 256
						},
						"rasp_id": {
							"type": "keyword",
							"ignore_above": 256
						},
						"event_time": {
							"type": "date"
						},
						"stack_trace": {
							"type": "keyword"
						},
						"intercept_state": {
							"type": "keyword",
							"ignore_above": 64
						},
						"attack_type": {
							"type": "keyword",
							"ignore_above": 256
						},
						"attack_location": {
							"type": "object",
							"properties": {
								"location_zh_cn":{
									"type": "keyword",
									"ignore_above": 256
								},
								"location_en":{
									"type": "keyword",
									"ignore_above": 256
								},
								"longitude":{
									"type": "double"
								},
								"latitude":{
									"type": "double"
								}
							}
						},
						"plugin_algorithm":{
							"type": "keyword",
							"ignore_above": 256
						},
						"plugin_name": {
							"type": "keyword",
							"ignore_above": 256
						},
						"plugin_confidence": {
							"type": "short"
						},
						"attack_params": {
							"type": "object",
							"enabled":"false"
						},
						"plugin_message": {
							"type": "keyword"
						},
						"server_nic": {
							"type": "nested",
							"properties": {
								"name": {
									"type": "keyword",
									"ignore_above": 256
								},
								"ip": {
									"type": "keyword",
									"ignore_above": 256
								}
							}
						}
					}
				}
			}
		}
	`
var policyAlarmTemplate = `
		{
			"template":"openrasp-policy-alarm-*",
			"aliases" : {
        		"real-{index}" : {} 
    		},
			"settings": {
				"analysis": {
					"normalizer": {
						"lowercase_normalizer": {
							"type": "custom",
							"filter": ["lowercase"]
						}
					}
				}
			},
			"mappings": {
				"policy-alarm": {
					"_all": {
						"enabled": false
					},
					"properties": {
						"@timestamp":{
							"type":"date"
						},
						"event_type": {
							"type": "keyword",
							"ignore_above": 256
						},
						"server_hostname": {
							"type": "keyword",
							"ignore_above": 256,
							"normalizer": "lowercase_normalizer"
						},
						"server_type": {
							"type": "keyword",
							"ignore_above": 64
						},
						"server_nic": {
							"type": "nested",
							"properties": {
								"name": {
									"type": "keyword",
									"ignore_above": 256
								},
								"ip": {
									"type": "keyword",
									"ignore_above": 256
								}
							}
						},
						"app_id": {
							"type": "keyword",
							"ignore_above": 256
						},
						"rasp_id": {
							"type": "keyword",
							"ignore_above": 256
						},
						"event_time": {
							"type": "date"
						},
						"stack_trace": {
							"type": "keyword"
						},
						"policy_id": {
							"type": "long"
						},
						"message": {
							"type": "keyword"
						},
						"stack_md5": {
							"type": "keyword",
							"ignore_above": 64
						},
						"policy_params": {
							"type": "object",
							"enabled":"false"
						}
					}
				}
			}
		}
		`
var errorAlarmTemplate = `{
			"template":"openrasp-error-alarm-*",
			"aliases" : {
        		"real-{index}" : {} 
    		},
			"settings": {
				"analysis": {
					"normalizer": {
						"lowercase_normalizer": {
							"type": "custom",
							"filter": ["lowercase"]
						}
					}     
				}
			},
			"mappings": {
				"error-alarm": {
					"_all": {
						"enabled": false
					},
					"properties": {
						"@timestamp":{
							"type":"date"
						},
						"app_id": {
							"type": "keyword",
							"ignore_above": 256
						},
						"rasp_id": {
							"type": "keyword",
							"ignore_above": 256
						},
						"message": {
							"type": "keyword"
						},
						"level": {
							"type": "keyword",
							"ignore_above": 64
						},
						"err_code":{
							"type": "long"
						},
						"stack_trace":{
							"type": "keyword"
						},
						"pid":{
							"type": "long"
						},
						"event_time": {
							"type": "date"
						},
						"server_hostname": {
							"type": "keyword",
							"ignore_above": 256,
							"normalizer": "lowercase_normalizer"
						},
						"server_nic": {
							"type": "nested",
							"properties": {
								"name": {
									"type": "keyword",
									"ignore_above": 256
								},
								"ip": {
									"type": "keyword",
									"ignore_above": 256
								}
							}
						}
					}
				}
			}
		}
	`
var reportDataTemplate = `
		{
			"template":"openrasp-report-data-*",
			"aliases" : {
        		"real-{index}" : {} 
    		},
			"mappings": {
				"report-data": {
					"_all": {
						"enabled": false
					},
					"properties": {
						"@timestamp":{
							"type":"date"
         				},
						"time": {
							"type": "date"
						},
						"request_sum": {
							"type": "long"
						},
						"rasp_id": {
							"type": "keyword",
							"ignore_above" : 256
						}
					}
				}
			}
		}
	`

func init() {
	CreateTemplate("report-data-template", reportDataTemplate)
	CreateTemplate("error-alarm-template", errorAlarmTemplate)
	CreateTemplate("attack-alarm-template", attackAlarmTemplate)
	CreateTemplate("policy-alarm-template", policyAlarmTemplate)
}
