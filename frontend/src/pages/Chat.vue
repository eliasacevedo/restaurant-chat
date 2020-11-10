<template>
    <div class="chat">
        <Profile :user="user" class="chat--profile"/>
        <v-container class="chat--messages" ref="chatContainer">
            <Message 
                v-for="message in messages" 
                :key="message.id"
                :time="message.time" 
                :origin="message.user" 
                :text="message.text" 
                :classification="message.classification"
            />
        </v-container> 
        <Input class="chat--input" @sendMessage="userAddMessage"/>
    </div>
     
</template>

<script lang="ts">
    import Message from '@/components/chat/Message.vue'
    import Profile from '@/components/chat/Profile.vue'
    import Input from '@/components/chat/Input.vue'
    import { IUserServices } from '@/services/users/users'
    import { Vue, Component, Inject } from 'vue-property-decorator'
    import { User } from '@/models/user'
    import { Message as IMessage } from '@/models/message'
    import { IBotServices } from '@/services/bots/bot'
    
    @Component({
        name: 'Chat',
        components: {
            Message,
            Profile,
            Input
        },
    })
    export default class Chat extends Vue{
        private user: User = { name: '', photoPath: '' };
        private userIdentifyOrigin = 'me';
        private botIdentifyOrigin = 'bot';
        private messageCount: number = 0;
        private messages: IMessage[] = []

        @Inject()
        userServices!: IUserServices;

        @Inject()
        botServices!: IBotServices;

        async mounted() {
            this.user = await this.userServices.getBotInfo();
        }

        addMessage(message: IMessage) {
            if (!message.id) {
                message.id = this.messageCount;
            }

            this.messages.push(message);

            const chatContainer = this.$refs.chatContainer as Element;

            if (!chatContainer) {
                return;
            }
            
            const heigth = chatContainer.scrollHeight;
            setTimeout(() => {
                chatContainer.scrollTo(0, heigth);
            }, 0)
            
            this.messageCount++
        }

        async userAddMessage(message: IMessage) {
            message.user = this.userIdentifyOrigin;
            this.addMessage(message);
            const response = await this.botServices.askQuestion(message);
            this.botAddMessage(response);
        }

        botAddMessage(message: IMessage) {
            message.user = this.botIdentifyOrigin;
            this.addMessage(message);
        }
    }
</script>

<style lang="scss" scoped>
    .chat {
        min-height: 100vh;
        display: grid;
        grid-template-rows: 65px 1fr 65px;
        width: 100%;

        &--messages {
            height: calc(100vh - 130px);
            overflow-x: hidden;
            overflow-y: auto;
        }
    }
</style>

