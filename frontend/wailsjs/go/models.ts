export namespace utils {
	
	export class User {
	    uid: string;
	    pwd: string;
	    profileID: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uid = source["uid"];
	        this.pwd = source["pwd"];
	        this.profileID = source["profileID"];
	        this.name = source["name"];
	    }
	}

}

