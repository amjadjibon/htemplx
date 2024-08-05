// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func PasteBin() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"pt-20 max-w-2xl mx-auto\"><div class=\"mb-5\"><label for=\"paste-content\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Paste Content</label> <textarea id=\"paste-content\" class=\"block w-full p-4 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-base focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" rows=\"6\"></textarea></div><div class=\"mb-5\"><label for=\"paste-name\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Name</label> <input type=\"text\" id=\"paste-name\" class=\"block w-full p-2.5 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-sm focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\"></div><div class=\"mb-5\"><label for=\"paste-category\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Category</label> <input type=\"text\" id=\"paste-category\" class=\"block w-full p-2.5 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-sm focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\"></div><div class=\"mb-5\"><label for=\"paste-tag\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Tag</label> <input type=\"text\" id=\"paste-tag\" class=\"block w-full p-2.5 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-sm focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\"></div><div class=\"mb-5\"><label class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Password Protection</label><div class=\"flex items-center mb-5\"><input id=\"password-enable\" type=\"checkbox\" name=\"password\" class=\"w-4 h-4 text-primary-700 border-gray-300 focus:ring-primary-700 dark:focus:ring-primary-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600\" _=\"on change if me.checked then set #password-field&#39;s style.display to &#39;block&#39; else set #password-field&#39;s style.display to &#39;none&#39;\"> <label for=\"password-enable\" class=\"block ml-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Enable Password</label></div><div id=\"password-field\" class=\"mb-5\" style=\"display: none;\"><label for=\"paste-password\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Password</label> <input type=\"password\" id=\"paste-password\" class=\"block w-full p-2.5 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-sm focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\"></div></div><div class=\"mb-5\"><label for=\"expiry-date-time\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Expiry Date & Time</label> <input type=\"datetime-local\" id=\"expiry-date-time\" class=\"block w-full p-2.5 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-sm focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\"></div><div><button type=\"submit\" class=\"py-3 px-5 w-full text-sm font-medium text-center text-white rounded-lg border cursor-pointer bg-primary-700 border-primary-600 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800\">Submit</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
