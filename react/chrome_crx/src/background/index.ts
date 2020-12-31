import { browser } from "webextension-polyfill-ts";

browser.runtime.onMessage.addListener((msg, sender) => {
    if (msg === "test - sendMessage and no listener answers") {
    }
    return Promise.resolve('true');
});