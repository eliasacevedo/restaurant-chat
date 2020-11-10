import { User } from '@/models/user';

export interface IUserServices {
    getActiveUser: () => Promise<User>;
    getBotInfo: () => Promise<User>;
}

export class UserServices implements IUserServices {

    private readonly ROUTES = {
        USER_INFO: 'http://localhost:9090/v1/user',
        BOT_INFO: 'http://localhost:9090/v1/user/bot'
    }

    async getActiveUser(): Promise<User> {
        const result = await fetch(this.ROUTES.USER_INFO)
        const data = await result.json() as User
        return data;
    }
    
    async getBotInfo(): Promise<User> {
        let data = {
            name: "",
            photoPath: ""
        };

        try {
            const result = await fetch(this.ROUTES.BOT_INFO);
            data = await result.json()
        } catch(e) {
            console.log(e, data)
        }
        
        return data;
    }
}

