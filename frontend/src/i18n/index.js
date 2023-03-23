
import {createI18n} from 'vue-i18n';
import zhLocale from './zh';

const messages = {
    zh :{
        ...zhLocale
    }
}

const i18n = createI18n({
    locale : localStorage.getItem('language') || 'zh',
    messages,
    fallbackLocale : 'zh',
    numberFormats:{
        'zh':{
            currency:{
                style:'currency',currency:'JPY',currencyDisplay:'symbol'
            }
        }
    },
    dateTimeFormats:{
        'zh':{
            short:{
                year:'numeric',month:'short',day:'numeric'
            },
            long:{
                year:'numeric',month:'short',day:'numeric',weekday:'short',hour:'numeric',minute:'numeric'
            }
        }
    }
})

export default i18n