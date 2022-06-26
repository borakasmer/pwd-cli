# pwd
With this CLI Tool, you can generate your own password with specific parameters. You can get custom password with "-s 7, -n 3, -c" flags.<br><br>
&#x1F34E;<I>This is Custom Password CLI. In every request, generated a new random password. If you don't set any parameters, you will get 8 length lowercase string password.</I>

<i>-c flag gets "true", if there is "-C" flag. If there is no -c flag, -c flag gets "false" value.<br> 
-s flag gets default 8 value.</i><br> <br>

<b><u>How to install Password Cli:</u></b><br>

<table><tr><td><img src="https://user-images.githubusercontent.com/9459881/165053981-38543faf-4bae-4500-8c28-fd5f497e0f46.gif"></img></td>
  <td><i>go install github.com/borakasmer/pwd@latest</i></td></tr></table>

<span style="color: red"><b>&#x1F534;Important:</b></span> You need Go program package to install pwd => <a href="https://go.dev/dl/" target="_blank">Go Downloads</a> </br>
<ol>
  <li>"go env" With the command "GOPATH" and "GOROOT" folders of GO can be seen.</li>
  <li>On <b>Mac</b> after "go install ..." the "go/bin/pwd" file under GOPATH => should be copied under "go/bin/" folder under GOROOT.</li>
  <li>In <b>Windows</b>, there is no need to further action :white_check_mark:</li>
</ol>
<img width="427" alt="FRJfHssXEAAqNsP" src="https://user-images.githubusercontent.com/9459881/165074359-572ca085-b1bd-4dbc-840f-43b1690a6319.png">
<b>:green_book: Extra Detail</b><br>
--------------------------------------------------------------------------------------------------------------
<ul>
  <li> "-s 10" The String length of Password is 10 and none of them is Capital. The default value is 8.</li>
  <li> "-n 7" => The Number length of Password</li>
  <li> "-s 5 -c" => The String Length of Password is 8 and some of them are Capital!</li>
  <li> "pwd" command => By default it means "pwd --str 8 --num 0". And there is no Capital letter in the Password.</li>
