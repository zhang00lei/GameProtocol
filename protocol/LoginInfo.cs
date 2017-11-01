//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: LoginInfo.proto
namespace PlayerLogin
{
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_PlayerLogin")]
  public partial class CS_PlayerLogin : global::ProtoBuf.IExtensible
  {
    public CS_PlayerLogin() {}
    
    private string _account = "";
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"account", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string account
    {
      get { return _account; }
      set { _account = value; }
    }
    private string _pwd = "";
    [global::ProtoBuf.ProtoMember(2, IsRequired = false, Name=@"pwd", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string pwd
    {
      get { return _pwd; }
      set { _pwd = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_PlayerLogin")]
  public partial class SC_PlayerLogin : global::ProtoBuf.IExtensible
  {
    public SC_PlayerLogin() {}
    
    private PlayerLogin.SC_PlayerLogin.LoginResult _loginResult = PlayerLogin.SC_PlayerLogin.LoginResult.SUCCESS;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"loginResult", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(PlayerLogin.SC_PlayerLogin.LoginResult.SUCCESS)]
    public PlayerLogin.SC_PlayerLogin.LoginResult loginResult
    {
      get { return _loginResult; }
      set { _loginResult = value; }
    }
    [global::ProtoBuf.ProtoContract(Name=@"LoginResult")]
    public enum LoginResult
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"SUCCESS", Value=0)]
      SUCCESS = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ACCOUNT_ERROR", Value=1)]
      ACCOUNT_ERROR = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"PWD_ERROR", Value=2)]
      PWD_ERROR = 2
    }
  
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_PlayerRegister")]
  public partial class CS_PlayerRegister : global::ProtoBuf.IExtensible
  {
    public CS_PlayerRegister() {}
    
    private string _account = "";
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"account", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string account
    {
      get { return _account; }
      set { _account = value; }
    }
    private string _pwd = "";
    [global::ProtoBuf.ProtoMember(2, IsRequired = false, Name=@"pwd", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string pwd
    {
      get { return _pwd; }
      set { _pwd = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_PlayerRegister")]
  public partial class SC_PlayerRegister : global::ProtoBuf.IExtensible
  {
    public SC_PlayerRegister() {}
    
    private PlayerLogin.SC_PlayerRegister.RegisterResult _registerResult = PlayerLogin.SC_PlayerRegister.RegisterResult.SUCCESS;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"registerResult", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(PlayerLogin.SC_PlayerRegister.RegisterResult.SUCCESS)]
    public PlayerLogin.SC_PlayerRegister.RegisterResult registerResult
    {
      get { return _registerResult; }
      set { _registerResult = value; }
    }
    [global::ProtoBuf.ProtoContract(Name=@"RegisterResult")]
    public enum RegisterResult
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"SUCCESS", Value=0)]
      SUCCESS = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ACCOUNT_ERROR", Value=1)]
      ACCOUNT_ERROR = 1
    }
  
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
}