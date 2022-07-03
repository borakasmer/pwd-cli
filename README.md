# pwd-cli
With this CLI Tool, you can generate your own password with specific parameters. You can get custom password with "-s 7, -n 3, -c -x" flags.<br><br>
&#x1F34E;<I>This is Custom Password CLI. In every request, generated a new random password. If you don't set any parameters, you will get 8 length lowercase string password.</I>

![NetflixGithub](https://raw.githubusercontent.com/borakasmer/pwd-cli/9f9277d27bda9e90da372c5035bde557ab0fa2d6/pwd-cli.gif)<br>

<b>Flags:</b>
<table><tr><td>:tv:</td><td><I>pwd-cli</I></td></tr><tr><td></td><td><I>pwd-cli -s 8 -n 3 -c -x</T></td></tr></table>
<ul>
  <li> -s, --str int        Length of string characters in Password</li>
  <li> -n, --num int        Length of number characters in Password</li>
  <li> -h, --help           Help for pwd-cli</li>
  <li> -c, --capital bool   Is there any capital letter ?</li>
  <li> -x, --symbol bool    Is there any symbol letter ? It's Length is allways => "len(str)/3 + 1"</li>
</ul>
<img width="600" alt="xflag" src="https://github.com/borakasmer/pwd-cli/blob/main/Screen%20Shot%202022-06-28%20at%2000.50.08.png">

<b>Default Command:</b> "pwd-cli" </br></br>
<b>Example Usage:</b>

<b>Usage:</b>
  pwd-cli [flags]
<ul>
  <li>"pwd-cli -s 8 -n 3 -c"</li>
  <li>"pwd-cli -s 10 -c -x"</li>
  <li>"pwd-cli -n 6"</li>
  <img alt="xflag" src="https://github.com/borakasmer/pwd-cli/blob/main/Screen%20Shot%202022-07-03%20at%2013.08.54.png">  
  <li>"pwd-cli --str 6 --num 2 --capital --symbol"</li>  
  <li>"pwd-cli" => default: 'pwd-cli -s 8 -n 0 -x 0 -c 0'</li>
</ul>
**********************************************************************************************************</br>

<i>-c flag gets "true", if there is "-c" flag. If there is no -c flag, -c flag gets "false" value.<br> 
-s flag gets default 8 value.<br> -x flag means put symbol in Password. Default symbol length formule is => "len(string)/3+1" </i><br> <br>

<b><u>How to install pwd-cli (Random Password Generator):</u></b><br>

<table><tr><td><img src="https://user-images.githubusercontent.com/9459881/165053981-38543faf-4bae-4500-8c28-fd5f497e0f46.gif"></img></td>
  <td><i>go install github.com/borakasmer/pwd-cli@latest</i></td></tr></table>

<span style="color: red"><b>&#x1F534;Important:</b></span> You need Go program package to install pwd-cli => <a href="https://go.dev/dl/" target="_blank">Go Downloads</a> </br>
<ol>
  <li>"go env" With the command "GOPATH" and "GOROOT" folders of GO can be seen.</li>
  <li>On <b>Mac</b> after "go install ..." the "go/bin/pwd-cli" file under GOPATH => should be copied under "go/bin/" folder under GOROOT.</li>
  <li>In <b>Windows</b>, there is no need to further action :white_check_mark:</li>
</ol>
<img width="427" alt="FRJfHssXEAAqNsP" src="https://user-images.githubusercontent.com/9459881/165074359-572ca085-b1bd-4dbc-840f-43b1690a6319.png">
<b>:green_book: Extra Detail</b><br>
--------------------------------------------------------------------------------------------------------------
<ul>
  <li> "-s 10" The Length of String Characters in Password is 10 and none of them is Capital. The default value is 8.</li>
  <li> "-n 7" => The Length of Number Characters in Password is 7</li>
  <li> "-s 5 -c" => The String Length of Password is 5 and some of them are Capital!</li>
   <li> "-s 6 -x" => The String Length of Password is 6 and (6/3+1) =>3 characters are Symbol! Total length of Password is 9</li>
  <li> "pwd-cli" command => By default it means "pwd-cli --str 8 --num 0". And there is no Capital letter in the Password.</li>
