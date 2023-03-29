export namespace models {
	
	export class GPM {
	    id: number;
	    website: string;
	    username: string;
	    email: string;
	    password: string;
	    notes: string;
	
	    static createFrom(source: any = {}) {
	        return new GPM(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.website = source["website"];
	        this.username = source["username"];
	        this.email = source["email"];
	        this.password = source["password"];
	        this.notes = source["notes"];
	    }
	}
	export class AllRows {
	    status: string;
	    rows: GPM[];
	
	    static createFrom(source: any = {}) {
	        return new AllRows(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.rows = this.convertValues(source["rows"], GPM);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class ReqInsertGPM {
	    website: string;
	    username: string;
	    email: string;
	    password: string;
	    notes: string;
	
	    static createFrom(source: any = {}) {
	        return new ReqInsertGPM(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.website = source["website"];
	        this.username = source["username"];
	        this.email = source["email"];
	        this.password = source["password"];
	        this.notes = source["notes"];
	    }
	}

}

