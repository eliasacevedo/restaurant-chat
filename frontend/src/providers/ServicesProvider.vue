<template>
    <div>
        <slot></slot>
    </div>
</template>

<script lang="ts">
import { Component, Provide, Vue } from 'vue-property-decorator';
import { BotServices, IBotServices } from '@/services/bots/bot';
import { IUserServices, UserServices } from '../services/users/users'
import { IUtility, UtilityServices } from '@/services/utilities/utility'

@Component({name: 'ServicesProvider'})
export default class ServicesProvider extends Vue{
    constructor() {
        super();
        this.userServices = new UserServices();
        this.botServices = new BotServices(this.userServices);
        this.utilityServices = new UtilityServices();
    }
    
    @Provide()
    userServices: IUserServices;

    @Provide()
    botServices: IBotServices;

    @Provide()
    utilityServices: IUtility

}
</script>

<style>

</style>
