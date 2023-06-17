export namespace fdu {
	
	export class Grade {
	    code: string;
	    year: string;
	    semester: string;
	    name: string;
	    credit: number;
	    grade: string;
	
	    static createFrom(source: any = {}) {
	        return new Grade(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.year = source["year"];
	        this.semester = source["semester"];
	        this.name = source["name"];
	        this.credit = source["credit"];
	        this.grade = source["grade"];
	    }
	}
	export class Grades {
	    grades: Grade[];
	    gpa: number;
	
	    static createFrom(source: any = {}) {
	        return new Grades(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.grades = this.convertValues(source["grades"], Grade);
	        this.gpa = source["gpa"];
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

}

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

