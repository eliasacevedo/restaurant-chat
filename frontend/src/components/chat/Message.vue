<template>
  <div class="message" >
      <div class="message--container" :data-origin="origin">
        <p class="message--container--text">{{text}}{{classification}}</p>
        <span class="message--container--time">{{Time}}</span>
      </div>
  </div>
</template>

<script lang="ts">
import { Component, Inject, Prop, Vue } from 'vue-property-decorator'
import { IUtility } from '@/services/utilities/utility'

@Component({ 
    name: 'Message',
})
export default class Message extends Vue{
    @Prop({required: true}) origin!: string
    @Prop({required: true}) text!: string
    @Prop({default: null}) classification!: string[]

    @Inject() 
    utilityServices!: IUtility

    @Prop({required: true}) time!: Date
    get Time(): string {
        return `${this.time.getHours()}:${this.utilityServices.addZeros(this.time.getMinutes(), 2)}`
    }
    
}

</script>

<style lang="scss" scoped>
    .message{
        width: 100%;
        position: relative;
        display: flex;
        flex-direction: column;
        margin-bottom: 5px;

        &--container {
            padding: 5px 10px 15px 10px;
            max-width: 60%;
            min-width: 50px;
            width: max-content;
            border-radius: 5px;
            position: relative;
            background-color: blue;
            color: white;

            &[data-origin]::before, &[data-origin]::after {
                position: absolute;
                width: 0;
                height: 0;
                border-style: solid;
                border-width: 5px 0 5px 20px;
            }

            &[data-origin="me"] {
                align-self: flex-end;                
                &::after {
                    right: -18px;
                    bottom: 0;
                    border-color: transparent transparent transparent blue;
                    content: '';
                }
            }

            &[data-origin="bot"] {
                color: black;
                background-color: lightgreen;
                align-self: flex-start;                
                &::before {
                    left: -18px;
                    top: 0;
                    transform: rotate(180deg);
                    border-color: transparent transparent transparent lightgreen;
                    content: '';
                }
            }

            &--text {
                
            }

            &--time {
                position: absolute;
                right: 5px;
                bottom: 2px;
                font-size: 0.7em;
            }
        }
    }
    
</style>
