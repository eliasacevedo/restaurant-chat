<template>
  <div class="form">
    <v-text-field
        outlined
        label="Mensaje"
        v-model="text"
        @keyup.enter="send"
        class="form--input"
        hide-details="auto"
    >
    </v-text-field>
    <v-btn icon large @click="send" type="submit" class="form--button">
        <v-icon color="blue">mdi-send</v-icon>
    </v-btn>
  </div>
</template>

<script lang="ts">

import { Message } from '@/models/message'
import { UserServices } from '@/services/users/users';
import {Component, Emit, Inject, Vue} from 'vue-property-decorator'

@Component({
    name: 'Input'
})
export default class Input extends Vue{

    @Inject()
    userServices!: UserServices;

    private text: string = '';
    @Emit('sendMessage')
    private async sendMessage(){

        const user = await this.userServices.getActiveUser();
        const message: Message = {
            text: this.text,
            user: user.name,
            time: new Date()
        }
        this.text = '';
        return message;
    }

    private async send() {
        if(!this.isValidMessage()) {
            return;
        }
        await this.sendMessage();
    }

    private isValidMessage(): boolean {
        if(this.text.trim().length === 0) {
            return false;
        }
        return true;
    }
}

</script>

<style lang="scss">
    .form{
        display: flex;
        align-items: center;
        padding: 5px;
        overflow: hidden;

        &--input{
            .v-text-field__details{
                display: none;
            }
        }

        &--button {
            background-color: blue;
        }
    }
</style>