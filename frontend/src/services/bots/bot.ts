import { Message } from '@/models/message';
import { User } from '@/models/user';
import { IUserServices } from '../users/users';

export interface IBotServices {
    askQuestion: (question: Message) => Promise<Message>;
}

export class BotServices implements IBotServices {
    
    private readonly ROUTES = {
        BOT_ASK: "http://localhost:9090/v1/bot/ask"
    }

    getBotInfo: () => Promise<User>;
    constructor({ getBotInfo }: IUserServices) {
        this.getBotInfo = getBotInfo;
    }

    async askQuestion(question: Message): Promise<Message> {
        const message = { ...question};
        const result = await fetch(this.ROUTES.BOT_ASK, {
            body: JSON.stringify(question),
            method: "POST"
        });
        const data = await result.json();
        message.text = data.text;
        message.classification = data.classification
        message.id = undefined
        return message;
    }
}



