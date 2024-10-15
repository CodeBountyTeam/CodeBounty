


interface Database {
/*
    The Database interface provides the interface for any database implementation. 
    It should only implement CRUD methods and should not be specific to certain data. 

    If a query fails, the implementation will throw an error 
*/



    create(tableName: string, data: object): Promise<void>;
    /* The Create method takes a TableName and inserts all data from the object into the database

	
    


    /* ---- FIX ME ---- */ 
    //get(tableName: string): Promise<T>;
    //getAll(): Promise<T[]>;
    //update(id: string, object: T): Promise<void>;
    //delete(id: string): Promise<void>;
}

export default Database

