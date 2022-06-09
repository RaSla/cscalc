namespace General_Client
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private async void button1_Click(object sender, EventArgs e)
        {
            using var client = new HttpClient();
            if (Uri.IsWellFormedUriString(textBox1.Text, UriKind.Absolute))
            {
                var check_for_valid_req = await client.SendAsync(new HttpRequestMessage(HttpMethod.Head, textBox1.Text));
                if (check_for_valid_req.ReasonPhrase == "OK")
                {
                    var req = await client.GetStringAsync(textBox1.Text);
                    label2.ForeColor = Color.Black;
                    label2.Text = "Result: \n\n" + req;
                }
                else
                {
                    label2.ForeColor = Color.Red;
                    label2.Text = "Request wasn't found on the server. Did you type the URL correctly?";
                }
            }
            else
            {
                label2.ForeColor = Color.Red;
                label2.Text = "Invalid URL!";
            }
        }
    }
}