<script setup lang="ts">
defineProps<{
  num: number
}>()
</script>

<template>
    <div class="poker-card">
        <span class="poker-number-top">
            <span class="poker-number">{{ num }}</span>
        </span>
        <span class="poker-number-bottom">
            <span class="poker-number">{{ num }}</span>
        </span>
    </div>
</template>
  
<style scoped>
    div.poker-card {
        border:1px solid var(--color-text);
        border-radius: .5em;
        min-width: 6em;
        max-width: 6em;
        min-height: 10em;
        max-height: 10em; 
        display: flex;
        flex-direction: column;
        margin:.5em;
        box-shadow: 0 3px 10px rgb(0 0 0 / 0.2);
        cursor: pointer;
        padding: .2em;

        span.poker-number {
            border: 1px solid var(--color-text);
            border-radius: 50%;
            display: flex;
            font-size: 1.4em;
            font-weight:bolder;
            margin: 0;
            padding: 0;
            width: 1.5em;
            height: 1.5em;
            align-items: center;
            text-align: center;
            justify-items: center;
            align-content: center;
            justify-content: center;
            background-color: var(--color-text);
            color: var(--vt-c-black);
        }
  
        span.poker-number-top {
            flex-grow: 1;
            align-self: flex-end;
            justify-self: end;
        }

        span.poker-number-bottom {
            flex-grow: 1;
            align-content: end;
        }
    }
  
    div.poker-card:hover {
        background-color:rgb(33, 61, 119);;
  
        span.poker-number {
            background-color: rgb(82, 124, 213);
            border-color:rgb(82, 124, 213);
        }
    }
</style>
  