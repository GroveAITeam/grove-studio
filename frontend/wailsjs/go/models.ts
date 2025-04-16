export namespace config {
	
	export class AppConfig {
	    version: string;
	    data_path: string;
	    debug_mode: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AppConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.data_path = source["data_path"];
	        this.debug_mode = source["debug_mode"];
	    }
	}

}

