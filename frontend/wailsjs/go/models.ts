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
	export class Rank {
	    current: number;
	    total: number;
	    percentage: number;
	    gpa: number;
	    credits: number;
	
	    static createFrom(source: any = {}) {
	        return new Rank(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current = source["current"];
	        this.total = source["total"];
	        this.percentage = source["percentage"];
	        this.gpa = source["gpa"];
	        this.credits = source["credits"];
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

export namespace xk {
	
	export class AmountInfo {
	    lc: number;
	    sc: number;
	
	    static createFrom(source: any = {}) {
	        return new AmountInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lc = source["lc"];
	        this.sc = source["sc"];
	    }
	}
	export class ArrangeInfo {
	    weekDay: number;
	    startUnit: number;
	    endUnit: number;
	
	    static createFrom(source: any = {}) {
	        return new ArrangeInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.weekDay = source["weekDay"];
	        this.startUnit = source["startUnit"];
	        this.endUnit = source["endUnit"];
	    }
	}
	export class Course {
	    name: string;
	    no: string;
	    code: string;
	    id: number;
	    courseId: number;
	    amount: AmountInfo;
	    arrangeInfo: ArrangeInfo[];
	    credits: number;
	    examTime: string;
	    examFormName: string;
	    teachers: string;
	    campusName: string;
	
	    static createFrom(source: any = {}) {
	        return new Course(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.no = source["no"];
	        this.code = source["code"];
	        this.id = source["id"];
	        this.courseId = source["courseId"];
	        this.amount = this.convertValues(source["amount"], AmountInfo);
	        this.arrangeInfo = this.convertValues(source["arrangeInfo"], ArrangeInfo);
	        this.credits = source["credits"];
	        this.examTime = source["examTime"];
	        this.examFormName = source["examFormName"];
	        this.teachers = source["teachers"];
	        this.campusName = source["campusName"];
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
	export class Query {
	    courseName: string;
	    lessonNo: string;
	    courseCode: string;
	
	    static createFrom(source: any = {}) {
	        return new Query(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.courseName = source["courseName"];
	        this.lessonNo = source["lessonNo"];
	        this.courseCode = source["courseCode"];
	    }
	}

}

