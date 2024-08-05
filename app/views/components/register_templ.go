// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Register() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"bg-gray-50 dark:bg-gray-900\"><div class=\"flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0\"><div class=\"w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700\"><div class=\"p-6 space-y-4 md:space-y-6 sm:p-8\"><h1 class=\"text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white\">Create an account</h1><form class=\"space-y-4 md:space-y-6\" action=\"#\"><div><label for=\"email\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Your email</label> <input type=\"email\" name=\"email\" id=\"email\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" placeholder=\"name@company.com\" required=\"\"></div><div><label for=\"password\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Password</label> <input type=\"password\" name=\"password\" id=\"password\" placeholder=\"••••••••\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" required=\"\"></div><div><label for=\"confirm-password\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Confirm password</label> <input type=\"confirm-password\" name=\"confirm-password\" id=\"confirm-password\" placeholder=\"••••••••\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" required=\"\"></div><div class=\"flex items-start\"><div class=\"flex items-center h-5\"><input id=\"terms\" aria-describedby=\"terms\" type=\"checkbox\" class=\"w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800\" required=\"\"></div><div class=\"ml-3 text-sm\"><label for=\"terms\" class=\"font-light text-gray-500 dark:text-gray-300\">I accept the <a class=\"font-medium text-primary-600 hover:underline dark:text-primary-500\" href=\"#\">Terms and Conditions</a></label></div></div><button hx-post=\"/sign-up\" hx-target=\"#login\" class=\"w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800\">Create an account</button> <button hx-get=\"/under-construction\" hx-target=\"#login\" class=\"w-full text-white bg-red-600 hover:bg-red-700 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-800\">Sign up with Google</button><div class=\"text-sm font-light text-gray-500 dark:text-gray-400\">Already have an account? <button hx-get=\"/login\" hx-target=\"#login\" hx-swap=\"innerHTML\" class=\"font-medium text-primary-600 hover:underline dark:text-primary-500\">Login here</button></div></form></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
