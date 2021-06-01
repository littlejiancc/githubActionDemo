REPLACE INTO `sp_dashboard_block_system` (`id`,`name`,`desc`,`scope`,`scope_id`,`version`,`view_config`) VALUES ("process_analysis_java","process_analysis_java","","micro_service","global","v2",'[{"h":10,"i":"view-gsdGSmsE","view":{"api":{"body":{"from":["docker_container_summary"],"groupby":["time()"],"select":[{"alias":"typeoQEtA1Qu","expr":"time()"},{"alias":"valuea5DMGiam","expr":"avg(cpu_usage_percent::field)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","filter_terminus_key":"{{terminusKey}}","format":"chartv2","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["docker_container_summary"],"typeDimensions":[{"alias":"时间","key":"typeoQEtA1Qu","type":"time"}],"valueDimensions":[{"alias":"core","key":"valuea5DMGiam","resultType":"number","type":"field"}]},"optionProps":{"invalidToZero":true,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":null,"title":"容器CPU"},"w":12,"x":0,"y":0},{"h":10,"i":"view-Vk0NhIvt","view":{"api":{"body":{"from":["docker_container_summary"],"groupby":["time()"],"select":[{"alias":"typeSf2rck4P","expr":"time()"},{"alias":"valuektRpqNyO","expr":"round_float(avg(mem_limit::field), 2)"},{"alias":"value9IeT8Ivv","expr":"round_float(avg(mem_usage::field), 2)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","filter__metric_scope":"micro_service","filter__metric_scope_id":"{{terminusKey}}","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","format":"chartv2","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["other","other@docker_container_summary"],"typeDimensions":[{"alias":"时间","key":"typeSf2rck4P","type":"time"}],"valueDimensions":[{"alias":"最大值","field":"docker_container_summary-mem_limit::field","key":"valuektRpqNyO","resultType":"number","type":"field","unit":{"type":"CAPACITY","unit":"B"}},{"alias":"已使用","field":"docker_container_summary-mem_usage::field","key":"value9IeT8Ivv","resultType":"number","type":"field","unit":{"type":"CAPACITY","unit":"B"}}]},"optionProps":{"invalidToZero":true,"isMoreThanOneDay":false,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":{},"title":"容器内存"},"w":12,"x":12,"y":0},{"h":10,"i":"view-U78akBad","view":{"api":{"body":{"from":["jvm_memory"],"groupby":["time()"],"select":[{"alias":"typeIt2X6865","expr":"time()"},{"alias":"valueKRHMRLK2","expr":"round_float(avg(used::field), 2)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","eq_name":"heap_memory","filter__metric_scope":"micro_service","filter__metric_scope_id":"{{terminusKey}}","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","format":"chartv2","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["application_jvm","application_jvm_memory","application_jvm_memory_heap"],"typeDimensions":[{"alias":"时间","key":"typeIt2X6865","type":"time"}],"valueDimensions":[{"alias":"已使用","field":"jvm_memory-used::field","key":"valueKRHMRLK2","resultType":"number","type":"field","unit":{"type":"CAPACITY","unit":"B"}}]},"optionProps":{"invalidToZero":true,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":null,"title":"Heap Memory Usage"},"w":12,"x":0,"y":10},{"h":10,"i":"view-lgddxxPr","view":{"api":{"body":{"from":["jvm_memory"],"groupby":["time()"],"select":[{"alias":"typeM0DAUMLM","expr":"time()"},{"alias":"value8nx5VIig","expr":"round_float(avg(used::field), 2)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","eq_name":"non_heap_memory","filter__metric_scope":"micro_service","filter__metric_scope_id":"{{terminusKey}}","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","format":"chartv2","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["application_jvm","application_jvm_memory","application_jvm_memory_non_heap"],"typeDimensions":[{"alias":"时间","key":"typeM0DAUMLM","type":"time"}],"valueDimensions":[{"alias":"已使用","field":"jvm_memory-used::field","key":"value8nx5VIig","resultType":"number","type":"field","unit":{"type":"CAPACITY","unit":"B"}}]},"optionProps":{"invalidToZero":true,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":null,"title":"Non Heap Memory Usage"},"w":12,"x":12,"y":10},{"h":10,"i":"view-246IruiQ","view":{"api":{"body":{"from":["jvm_memory"],"groupby":["time()"],"select":[{"alias":"typePnIyzrBO","expr":"time()"},{"alias":"valueIXILX6JH","expr":"round_float(avg(used::field), 2)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","filter__metric_scope":"micro_service","filter__metric_scope_id":"{{terminusKey}}","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","format":"chartv2","match_name":"*_eden_space","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["application_jvm","application_jvm_memory","application_jvm_memory_ps_eden_space"],"typeDimensions":[{"alias":"时间","key":"typePnIyzrBO","type":"time"}],"valueDimensions":[{"alias":"已使用","field":"jvm_memory-used::field","key":"valueIXILX6JH","resultType":"number","type":"field","unit":{"type":"CAPACITY","unit":"B"}}]},"optionProps":{"invalidToZero":true,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":null,"title":"Eden Space"},"w":8,"x":0,"y":20},{"h":10,"i":"view-AzNTzbgm","view":{"api":{"body":{"from":["jvm_memory"],"groupby":["time()"],"select":[{"alias":"typeqDJ0hjqO","expr":"time()"},{"alias":"valueNQr0xQ5n","expr":"round_float(avg(used::field), 2)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","filter__metric_scope":"micro_service","filter__metric_scope_id":"{{terminusKey}}","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","format":"chartv2","match_name":"*_old_gen","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["application_jvm","application_jvm_memory","application_jvm_memory_ps_old_gen"],"typeDimensions":[{"alias":"时间","key":"typeqDJ0hjqO","type":"time"}],"valueDimensions":[{"alias":"已使用","field":"jvm_memory-used::field","key":"valueNQr0xQ5n","resultType":"number","type":"field","unit":{"type":"CAPACITY","unit":"B"}}]},"optionProps":{"invalidToZero":true,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":null,"title":"Old Gen"},"w":8,"x":8,"y":20},{"h":10,"i":"view-gXyQwxmE","view":{"api":{"body":{"from":["jvm_memory"],"groupby":["time()"],"select":[{"alias":"typeWqvqa3pT","expr":"time()"},{"alias":"value8pbaaXGT","expr":"round_float(avg(used::field), 2)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","filter__metric_scope":"micro_service","filter__metric_scope_id":"{{terminusKey}}","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","format":"chartv2","match_name":"*_survivor_space","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["application_jvm","application_jvm_memory","application_jvm_memory_ps_surv_space"],"typeDimensions":[{"alias":"时间","key":"typeWqvqa3pT","type":"time"}],"valueDimensions":[{"alias":"已使用","field":"jvm_memory-used::field","key":"value8pbaaXGT","resultType":"number","type":"field","unit":{"type":"CAPACITY","unit":"B"}}]},"optionProps":{"invalidToZero":true,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":null,"title":"Survivor Space"},"w":8,"x":16,"y":20},{"h":10,"i":"view-BMxjOM9K","view":{"api":{"body":{"from":["jvm_class_loader"],"groupby":["time()"],"select":[{"alias":"typefmoF7423","expr":"time()"},{"alias":"value9xCtZWjF","expr":"round_float(avg(loaded::field), 2)"},{"alias":"unloaded","expr":"round_float(avg(unloaded::field), 2)"}]},"extraData":null,"header":null,"method":"post","query":{"end":"{{endTime}}","epoch":"ms","filter__metric_scope":"micro_service","filter__metric_scope_id":"{{terminusKey}}","filter_service_id":"{{serviceId}}","filter_service_instance_id":"{{instanceId}}","filter_service_name":"{{serviceName}}","format":"chartv2","ql":"influxql:ast","start":"{{startTime}}","type":"_"},"url":"/api/tmc/metrics-query"},"chartType":"chart:line","config":{"dataSourceConfig":{"activedMetricGroups":["jvm_class_loader"],"typeDimensions":[{"alias":"时间","key":"typefmoF7423","type":"time"}],"valueDimensions":[{"alias":"loaded","key":"value9xCtZWjF","resultType":"number","type":"field"},{"alias":"unloaded","key":"unloaded","resultType":"number","type":"field"}]},"optionProps":{"invalidToZero":true,"noAreaColor":true}},"controls":null,"dataSourceType":"api","description":"","staticData":null,"title":"Class Count"},"w":12,"x":0,"y":30},{"h":10,"i":"view-xRTxHcRk","view":{"api":{"body":{"from":["jjkkkjjj
