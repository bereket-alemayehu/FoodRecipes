import { defineNuxtPlugin } from "#app";
import { defineRule, configure } from "vee-validate";
import { required, email, min } from "@vee-validate/rules";

export default defineNuxtPlugin((nuxtApp) => {
  defineRule("required", required);
  defineRule("email", email);
  defineRule("min", min);

  configure({
    generateMessage: ({ rule, field }) => {
      const messages = {
        required: `${field} is required.`,
        email: `${field} must be a valid email.`,
        min: `${field} must be at least ${rule.params[0]} characters.`,
      };
      return messages[rule.name];
    },
  });
});
