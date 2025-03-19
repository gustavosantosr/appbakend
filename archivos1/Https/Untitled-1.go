


public static CreationDocumentRequest CanonicalRequestToCreationDocumentRequest(Core.Model.Pagares pagares)
{
    var documents = new List<Documents> { };
    var templateFields = new TemplateFields { };
    var notificationfrequency = new notificationFrequency { };

    // Actualización de DocumentSignatures para cumplir con el formato requerido
    var documentSignatures = new Dictionary<string, List<Signature>>
{
    {
        "1", new List<Signature>
        {
            new Signature
            {
                type = SignatureType.NAME_AND_DESCRIPTION
            }
        }
    }
};
    var signers = new List<Signer> { };

    string TipoPagare = pagares.Pagare.TipoPagare;

    // Busca el valor del enumerador basándote en la descripción
    var enumValue = Enum.GetValues(typeof(Templates))
                        .Cast<Templates>()
                        .FirstOrDefault(e => e.GetDescription() == TipoPagare);
    switch ((int)enumValue)
    {
        case 0:
        case 3:
        case 19:
            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? (pagares.Pagare.Persona[2].DatosContacto.NumeroCelular == 0)
? ""
: pagares.Pagare.Persona[2].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[1].DatosContacto.NumeroCelular == 0)
? ""
: pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = pagares.Pagare.Persona[1].DatosContacto.Email,
                NUMPOLIZA = pagares.Pagare.Poliza.Numero,
                NSUCNRAMNPOL = pagares.Pagare.Sucursal.Codigo + "-" + pagares.Pagare.Ramo.Codigo + "-" + pagares.Pagare.Poliza.Numero


            };
            signers = new List<Signer>
                    {
                        new Signer
                        {
                            ExternalUser = true,
                            Email =  pagares.Pagare.Persona[1].DatosContacto.Email,
                            FirstName = pagares.Pagare.Persona[1].Nombre1,
                            LastName = pagares.Pagare.Persona[1].Apellido1,
                            DocumentId =  pagares.Pagare.Persona[1].Documento.Numero,
                            DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                            MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                            CertificateProfile = "External Electronic",
                            //Validation = ValidationType.NONE,
                           // OtpTypes = new List<OtpType> { OtpType.SMS, OtpType.EMAIL }
                        }
                    };

            documents = new List<Documents>
                    {
                        new Documents
                         {

                                TemplateId = TipoPagare,
                                TemplateFields = templateFields,
                                DocumentName = pagares.Pagare.Id,
                                Main = true,
                                SignatureRequired = true,




                         },
                        new Documents
                         {

                                TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                Main = false,
                                SignatureRequired = true,




                         },


                    };

            break;
        case 1:
        case 2:
        case 4:
        case 5:
            signers = new List<Signer>();
            if (pagares.Pagare.Persona.Count > 1 && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[1].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[1].Nombre1,
                    LastName = pagares.Pagare.Persona[1].Apellido1,
                    DocumentId = pagares.Pagare.Persona[1].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 3 && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[3].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[3].Nombre1,
                    LastName = pagares.Pagare.Persona[3].Apellido1,
                    DocumentId = pagares.Pagare.Persona[3].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 5 && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[5].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[5].Nombre1,
                    LastName = pagares.Pagare.Persona[5].Apellido1,
                    DocumentId = pagares.Pagare.Persona[5].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 7 && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[7].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[7].Nombre1,
                    LastName = pagares.Pagare.Persona[7].Apellido1,
                    DocumentId = pagares.Pagare.Persona[7].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }

            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = pagares.Pagare.Persona[1].DatosContacto.Email,
                NUMPOLIZA = pagares.Pagare.Poliza.Numero,
                NSUCNRAMNPOL = pagares.Pagare.Sucursal.Codigo + "-" + pagares.Pagare.Ramo.Codigo + "-" + pagares.Pagare.Poliza.Numero,

                // Asignación de valores para persona[3]
                NOMAPCOD1 = pagares.Pagare.Persona[3].Nombre1 + ' ' + pagares.Pagare.Persona[3].Nombre2 + ' ' + pagares.Pagare.Persona[3].Apellido1 + ' ' + pagares.Pagare.Persona[3].Apellido2,
                TIPODOCCOD1 = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                NUMIDCOD1 = pagares.Pagare.Persona[3].Documento.Numero,
                DIRCOD1 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[4].Ubicacion.Direccion) ? pagares.Pagare.Persona[4].Ubicacion.Direccion : pagares.Pagare.Persona[3].Ubicacion.Direccion,
                TELCOD1 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? (pagares.Pagare.Persona[4].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[4].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[3].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].Ubicacion.Ciudad : pagares.Pagare.Persona[3].Ubicacion.Ciudad,
                CORREOCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].DatosContacto.Email : pagares.Pagare.Persona[3].DatosContacto.Email,

                // Asignación de valores para persona[4]

                RAZONSOCCOD1 = pagares.Pagare.Persona[4].Nombre1 + ' ' + pagares.Pagare.Persona[4].Nombre2 + ' ' + pagares.Pagare.Persona[4].Apellido1 + ' ' + pagares.Pagare.Persona[4].Apellido2,
                TIPODOCRSCOD1 = pagares.Pagare.Persona[4].Documento.TipoDocumento.Id,
                NUMIDRSCOD1 = pagares.Pagare.Persona[4].Documento.Numero,

                // Asignación de valores para persona[5]
                NOMAPCOD2 = pagares.Pagare.Persona[5].Nombre1 + ' ' + pagares.Pagare.Persona[5].Nombre2 + ' ' + pagares.Pagare.Persona[5].Apellido1 + ' ' + pagares.Pagare.Persona[5].Apellido2,
                TIPODOCCOD2 = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                NUMIDCOD2 = pagares.Pagare.Persona[5].Documento.Numero,
                DIRCOD2 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[6].Ubicacion.Direccion) ? pagares.Pagare.Persona[6].Ubicacion.Direccion : pagares.Pagare.Persona[5].Ubicacion.Direccion,
                TELCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? (pagares.Pagare.Persona[6].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[6].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[5].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].Ubicacion.Ciudad : pagares.Pagare.Persona[5].Ubicacion.Ciudad,
                CORREOCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].DatosContacto.Email : pagares.Pagare.Persona[5].DatosContacto.Email,


                // Asignación de valores para persona[6]
                RAZONSOCCOD2 = pagares.Pagare.Persona[6].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD2 = pagares.Pagare.Persona[6].Documento.TipoDocumento.Id,
                NUMIDRSCOD2 = pagares.Pagare.Persona[6].Documento.Numero,

                // Asignación de valores para persona[7]
                NOMAPCOD3 = pagares.Pagare.Persona[7].Nombre1 + ' ' + pagares.Pagare.Persona[7].Nombre2 + ' ' + pagares.Pagare.Persona[7].Apellido1 + ' ' + pagares.Pagare.Persona[7].Apellido2,
                TIPODOCCOD3 = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                NUMIDCOD3 = pagares.Pagare.Persona[7].Documento.Numero,
                DIRCOD3 = pagares.Pagare.Persona[7].Ubicacion.Direccion,
                TELCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? (pagares.Pagare.Persona[8].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[8].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[7].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].Ubicacion.Ciudad : pagares.Pagare.Persona[7].Ubicacion.Ciudad,
                CORREOCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].DatosContacto.Email : pagares.Pagare.Persona[7].DatosContacto.Email,

                // Asignación de valores para persona[7]
                RAZONSOCCOD3 = pagares.Pagare.Persona[8].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD3 = pagares.Pagare.Persona[8].Documento.TipoDocumento.Id,
                NUMIDRSCOD3 = pagares.Pagare.Persona[8].Documento.Numero,

            };
            documents = new List<Documents>
                {
                    new Documents
                     {

                            TemplateId = TipoPagare,
                            TemplateFields = templateFields,
                            DocumentName = pagares.Pagare.Id,
                            Main = true,
                            SignatureRequired = true,




                     },
                    new Documents
                     {

                            TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                            DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                            Main = false,
                            SignatureRequired = true,




                     },


                };

            break;
        case 14:
        case 16:
        case 18:
            signers = new List<Signer>();
            if (pagares.Pagare.Persona.Count > 1 && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[1].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[1].Nombre1,
                    LastName = pagares.Pagare.Persona[1].Apellido1,
                    DocumentId = pagares.Pagare.Persona[1].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 9 && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[9].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[9].Nombre1,
                    LastName = pagares.Pagare.Persona[9].Apellido1,
                    DocumentId = pagares.Pagare.Persona[9].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 11 && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[11].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[11].Nombre1,
                    LastName = pagares.Pagare.Persona[11].Apellido1,
                    DocumentId = pagares.Pagare.Persona[11].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 13 && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[13].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[13].Nombre1,
                    LastName = pagares.Pagare.Persona[13].Apellido1,
                    DocumentId = pagares.Pagare.Persona[13].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? (pagares.Pagare.Persona[2].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[2].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[1].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].Ubicacion.Ciudad : pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.Email : pagares.Pagare.Persona[1].DatosContacto.Email,

                RAZONSOCTOM = pagares.Pagare.Persona[2].Nombre1 + ' ' + pagares.Pagare.Persona[2].Nombre2 + ' ' + pagares.Pagare.Persona[2].Apellido1 + ' ' + pagares.Pagare.Persona[2].Apellido2,
                TIPODOCTOMPJ = pagares.Pagare.Persona[2].Documento.TipoDocumento.Id,
                NUMIDTOMPJ = pagares.Pagare.Persona[2].Documento.Numero,


                // Asignación de valores para persona[9]
                NOMAPINT1 = pagares.Pagare.Persona[9].Nombre1 + ' ' + pagares.Pagare.Persona[9].Nombre2 + ' ' + pagares.Pagare.Persona[9].Apellido1 + ' ' + pagares.Pagare.Persona[9].Apellido2,
                TIPODOCINT1 = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                NUMIDINT1 = pagares.Pagare.Persona[9].Documento.Numero,
                DIRINT1 = pagares.Pagare.Persona[9].Ubicacion.Direccion,
                TELINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? (pagares.Pagare.Persona[10].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[10].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[9].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].Ubicacion.Ciudad : pagares.Pagare.Persona[9].Ubicacion.Ciudad,
                CORREOINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].DatosContacto.Email : pagares.Pagare.Persona[9].DatosContacto.Email,

                // Asignación de valores para persona[10]
                RAZONSOCINT1 = pagares.Pagare.Persona[10].Nombre1 + ' ' + pagares.Pagare.Persona[10].Nombre2 + ' ' + pagares.Pagare.Persona[10].Apellido1 + ' ' + pagares.Pagare.Persona[10].Apellido2,
                TIPODOCCONUNINT1 = pagares.Pagare.Persona[10].Documento.TipoDocumento.Id,
                NUMIDCONUNINT1 = pagares.Pagare.Persona[10].Documento.Numero,

                // Asignación de valores para persona[11]
                NOMAPINT2 = pagares.Pagare.Persona[11].Nombre1 + ' ' + pagares.Pagare.Persona[11].Nombre2 + ' ' + pagares.Pagare.Persona[11].Apellido1 + ' ' + pagares.Pagare.Persona[11].Apellido2,
                TIPODOCINT2 = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                NUMIDINT2 = pagares.Pagare.Persona[11].Documento.Numero,
                DIRINT2 = pagares.Pagare.Persona[11].Ubicacion.Direccion,
                TELINT2 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? (pagares.Pagare.Persona[12].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[12].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[11].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].Ubicacion.Ciudad : pagares.Pagare.Persona[11].Ubicacion.Ciudad,
                CORREOINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].DatosContacto.Email : pagares.Pagare.Persona[11].DatosContacto.Email,


                // Asignación de valores para persona[12]
                RAZONSOCINT2 = pagares.Pagare.Persona[12].Nombre1 + ' ' + pagares.Pagare.Persona[12].Nombre2 + ' ' + pagares.Pagare.Persona[12].Apellido1 + ' ' + pagares.Pagare.Persona[12].Apellido2,
                TIPODOCCONUNINT2 = pagares.Pagare.Persona[12].Documento.TipoDocumento.Id,
                NUMIDCONUNINT2 = pagares.Pagare.Persona[12].Documento.Numero,

                // Asignación de valores para persona[13]
                NOMAPINT3 = pagares.Pagare.Persona[13].Nombre1 + ' ' + pagares.Pagare.Persona[13].Nombre2 + ' ' + pagares.Pagare.Persona[13].Apellido1 + ' ' + pagares.Pagare.Persona[13].Apellido2,
                TIPODOCINT3 = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                NUMIDINT3 = pagares.Pagare.Persona[13].Documento.Numero,
                DIRINT3 = pagares.Pagare.Persona[13].Ubicacion.Direccion,
                TELINT3 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? (pagares.Pagare.Persona[14].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[14].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[13].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].Ubicacion.Ciudad : pagares.Pagare.Persona[13].Ubicacion.Ciudad,
                CORREOINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].DatosContacto.Email : pagares.Pagare.Persona[13].DatosContacto.Email,

                // Asignación de valores para persona[14]
                RAZONSOCINT3 = pagares.Pagare.Persona[14].Nombre1 + ' ' + pagares.Pagare.Persona[14].Nombre2 + ' ' + pagares.Pagare.Persona[14].Apellido1 + ' ' + pagares.Pagare.Persona[14].Apellido2,
                TIPODOCCONUNINT3 = pagares.Pagare.Persona[14].Documento.TipoDocumento.Id,
                NUMIDCONUNINT3 = pagares.Pagare.Persona[14].Documento.Numero,


            };


            documents = new List<Documents>
                    {
                        new Documents
                         {

                                TemplateId = TipoPagare,
                                TemplateFields = templateFields,
                                DocumentName = pagares.Pagare.Id,
                                Main = true,
                                SignatureRequired = true,




                         },
                        new Documents
                         {

                                TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                Main = false,
                                SignatureRequired = true,




                         },


                    };
            break;
        case 7:
        case 9:
        case 11:
        case 13:
            signers = new List<Signer>();
            if (pagares.Pagare.Persona.Count > 1 && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[1].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[1].Nombre1,
                    LastName = pagares.Pagare.Persona[1].Apellido1,
                    DocumentId = pagares.Pagare.Persona[1].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 3 && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[3].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[3].Nombre1,
                    LastName = pagares.Pagare.Persona[3].Apellido1,
                    DocumentId = pagares.Pagare.Persona[3].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 5 && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[5].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[5].Nombre1,
                    LastName = pagares.Pagare.Persona[5].Apellido1,
                    DocumentId = pagares.Pagare.Persona[5].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 7 && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[7].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[7].Nombre1,
                    LastName = pagares.Pagare.Persona[7].Apellido1,
                    DocumentId = pagares.Pagare.Persona[7].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }

            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? (pagares.Pagare.Persona[2].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[2].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[1].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].Ubicacion.Ciudad : pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.Email : pagares.Pagare.Persona[1].DatosContacto.Email,
                NSUCNRAMNPOL = pagares.Pagare.Sucursal.Codigo + "-" + pagares.Pagare.Ramo.Codigo + "-" + pagares.Pagare.Poliza.Numero,
                RAZONSOCTOM = pagares.Pagare.Persona[2].Nombre1 + ' ' + pagares.Pagare.Persona[2].Nombre2 + ' ' + pagares.Pagare.Persona[2].Apellido1 + ' ' + pagares.Pagare.Persona[2].Apellido2,
                TIPODOCTOMPJ = pagares.Pagare.Persona[2].Documento.TipoDocumento.Id,
                NUMIDTOMPJ = pagares.Pagare.Persona[2].Documento.Numero,
                NUMPOLIZA = pagares.Pagare.Poliza.Numero,

                // Asignación de valores para persona[3]
                NOMAPCOD1 = pagares.Pagare.Persona[3].Nombre1 + ' ' + pagares.Pagare.Persona[3].Nombre2 + ' ' + pagares.Pagare.Persona[3].Apellido1 + ' ' + pagares.Pagare.Persona[3].Apellido2,
                TIPODOCCOD1 = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                NUMIDCOD1 = pagares.Pagare.Persona[3].Documento.Numero,
                DIRCOD1 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[4].Ubicacion.Direccion) ? pagares.Pagare.Persona[4].Ubicacion.Direccion : pagares.Pagare.Persona[3].Ubicacion.Direccion,



                TELCOD1  = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? (pagares.Pagare.Persona[4].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[4].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[3].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].Ubicacion.Ciudad : pagares.Pagare.Persona[3].Ubicacion.Ciudad,
                CORREOCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].DatosContacto.Email : pagares.Pagare.Persona[3].DatosContacto.Email,



                // Asignación de valores para persona[4]

                RAZONSOCCOD1 = pagares.Pagare.Persona[4].Nombre1 + ' ' + pagares.Pagare.Persona[4].Nombre2 + ' ' + pagares.Pagare.Persona[4].Apellido1 + ' ' + pagares.Pagare.Persona[4].Apellido2,
                TIPODOCRSCOD1 = pagares.Pagare.Persona[4].Documento.TipoDocumento.Id,
                NUMIDRSCOD1 = pagares.Pagare.Persona[4].Documento.Numero,

                // Asignación de valores para persona[5]
                NOMAPCOD2 = pagares.Pagare.Persona[5].Nombre1 + ' ' + pagares.Pagare.Persona[5].Nombre2 + ' ' + pagares.Pagare.Persona[5].Apellido1 + ' ' + pagares.Pagare.Persona[5].Apellido2,
                TIPODOCCOD2 = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                NUMIDCOD2 = pagares.Pagare.Persona[5].Documento.Numero,
                DIRCOD2 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[6].Ubicacion.Direccion) ? pagares.Pagare.Persona[6].Ubicacion.Direccion : pagares.Pagare.Persona[5].Ubicacion.Direccion,
                TELCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? (pagares.Pagare.Persona[6].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[6].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[5].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].Ubicacion.Ciudad : pagares.Pagare.Persona[5].Ubicacion.Ciudad,
                CORREOCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].DatosContacto.Email : pagares.Pagare.Persona[5].DatosContacto.Email,


                // Asignación de valores para persona[6]
                RAZONSOCCOD2 = pagares.Pagare.Persona[6].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD2 = pagares.Pagare.Persona[6].Documento.TipoDocumento.Id,
                NUMIDRSCOD2 = pagares.Pagare.Persona[6].Documento.Numero,

                // Asignación de valores para persona[7]
                NOMAPCOD3 = pagares.Pagare.Persona[7].Nombre1 + ' ' + pagares.Pagare.Persona[7].Nombre2 + ' ' + pagares.Pagare.Persona[7].Apellido1 + ' ' + pagares.Pagare.Persona[7].Apellido2,
                TIPODOCCOD3 = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                NUMIDCOD3 = pagares.Pagare.Persona[7].Documento.Numero,
                DIRCOD3 = pagares.Pagare.Persona[7].Ubicacion.Direccion,
                TELCOD3  = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? (pagares.Pagare.Persona[8].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[8].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[7].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].Ubicacion.Ciudad : pagares.Pagare.Persona[7].Ubicacion.Ciudad,
                CORREOCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].DatosContacto.Email : pagares.Pagare.Persona[7].DatosContacto.Email,

                // Asignación de valores para persona[7]
                RAZONSOCCOD3 = pagares.Pagare.Persona[8].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD3 = pagares.Pagare.Persona[8].Documento.TipoDocumento.Id,
                NUMIDRSCOD3 = pagares.Pagare.Persona[8].Documento.Numero,

            };
            documents = new List<Documents>
                {
                    new Documents
                     {

                            TemplateId = TipoPagare,
                            TemplateFields = templateFields,
                            DocumentName = pagares.Pagare.Id,
                            Main = true,
                            SignatureRequired = true,




                     },
                    new Documents
                     {

                            TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                            DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                            Main = false,
                            SignatureRequired = true



                     },


                };
            break;
        case 6:
        case 8:
        case 10:
        case 12:
            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM  = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? (pagares.Pagare.Persona[2].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[2].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[1].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].Ubicacion.Ciudad : pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.Email : pagares.Pagare.Persona[1].DatosContacto.Email,
                RAZONSOCTOM = pagares.Pagare.Persona[2].Nombre1 + ' ' + pagares.Pagare.Persona[2].Nombre2 + ' ' + pagares.Pagare.Persona[2].Apellido1 + ' ' + pagares.Pagare.Persona[2].Apellido2,
                TIPODOCTOMPJ = pagares.Pagare.Persona[2].Documento.TipoDocumento.Id,
                NUMIDTOMPJ = pagares.Pagare.Persona[2].Documento.Numero,
                NSUCNRAMNPOL = pagares.Pagare.Sucursal.Codigo + "-" + pagares.Pagare.Ramo.Codigo + "-" + pagares.Pagare.Poliza.Numero
            };
            signers = new List<Signer>
                    {
                        new Signer
                        {
                            ExternalUser = true,
                            Email =  pagares.Pagare.Persona[1].DatosContacto.Email,
                            FirstName = pagares.Pagare.Persona[1].Nombre1,
                            LastName = pagares.Pagare.Persona[1].Apellido1,
                            DocumentId =  pagares.Pagare.Persona[1].Documento.Numero,
                            DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                            MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                            CertificateProfile = "External Electronic",
                            //Validation = ValidationType.NONE,
                           // OtpTypes = new List<OtpType> { OtpType.SMS, OtpType.EMAIL }
                        }
                    };

            documents = new List<Documents>
                    {
                        new Documents
                         {

                                TemplateId = TipoPagare,
                                TemplateFields = templateFields,
                                DocumentName = pagares.Pagare.Id,
                                Main = true,
                                SignatureRequired = true,




                         },
                        new Documents
                         {

                                TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                Main = false,
                                SignatureRequired = true,




                         },


                    };
            break;
        case 15:
        case 17:
            signers = new List<Signer>();
            if (pagares.Pagare.Persona.Count > 1 && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[1].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[1].Nombre1,
                    LastName = pagares.Pagare.Persona[1].Apellido1,
                    DocumentId = pagares.Pagare.Persona[1].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 3 && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[3].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[3].Nombre1,
                    LastName = pagares.Pagare.Persona[3].Apellido1,
                    DocumentId = pagares.Pagare.Persona[3].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 5 && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[5].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[5].Nombre1,
                    LastName = pagares.Pagare.Persona[5].Apellido1,
                    DocumentId = pagares.Pagare.Persona[5].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 7 && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[7].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[7].Nombre1,
                    LastName = pagares.Pagare.Persona[7].Apellido1,
                    DocumentId = pagares.Pagare.Persona[7].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 9 && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[9].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[9].Nombre1,
                    LastName = pagares.Pagare.Persona[9].Apellido1,
                    DocumentId = pagares.Pagare.Persona[9].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 11 && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[11].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[11].Nombre1,
                    LastName = pagares.Pagare.Persona[11].Apellido1,
                    DocumentId = pagares.Pagare.Persona[11].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 13 && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[13].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[13].Nombre1,
                    LastName = pagares.Pagare.Persona[13].Apellido1,
                    DocumentId = pagares.Pagare.Persona[13].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM  = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? (pagares.Pagare.Persona[2].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[2].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[1].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].Ubicacion.Ciudad : pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.Email : pagares.Pagare.Persona[1].DatosContacto.Email,
                RAZONSOCTOM = pagares.Pagare.Persona[2].Nombre1 + ' ' + pagares.Pagare.Persona[2].Nombre2 + ' ' + pagares.Pagare.Persona[2].Apellido1 + ' ' + pagares.Pagare.Persona[2].Apellido2,
                TIPODOCTOMPJ = pagares.Pagare.Persona[2].Documento.TipoDocumento.Id,
                NUMIDTOMPJ = pagares.Pagare.Persona[2].Documento.Numero,


                // Asignación de valores para persona[3]
                NOMAPCOD1 = pagares.Pagare.Persona[3].Nombre1 + ' ' + pagares.Pagare.Persona[3].Nombre2 + ' ' + pagares.Pagare.Persona[3].Apellido1 + ' ' + pagares.Pagare.Persona[3].Apellido2,
                TIPODOCCOD1 = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                NUMIDCOD1 = pagares.Pagare.Persona[3].Documento.Numero,
                DIRCOD1 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[4].Ubicacion.Direccion) ? pagares.Pagare.Persona[4].Ubicacion.Direccion : pagares.Pagare.Persona[3].Ubicacion.Direccion,
                TELCOD1  = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? (pagares.Pagare.Persona[4].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[4].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[3].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].Ubicacion.Ciudad : pagares.Pagare.Persona[3].Ubicacion.Ciudad,
                CORREOCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].DatosContacto.Email : pagares.Pagare.Persona[3].DatosContacto.Email,
                // Asignación de valores para persona[4]

                RAZONSOCCOD1 = pagares.Pagare.Persona[4].Nombre1 + ' ' + pagares.Pagare.Persona[4].Nombre2 + ' ' + pagares.Pagare.Persona[4].Apellido1 + ' ' + pagares.Pagare.Persona[4].Apellido2,
                TIPODOCRSCOD1 = pagares.Pagare.Persona[4].Documento.TipoDocumento.Id,
                NUMIDRSCOD1 = pagares.Pagare.Persona[4].Documento.Numero,

                // Asignación de valores para persona[5]
                NOMAPCOD2 = pagares.Pagare.Persona[5].Nombre1 + ' ' + pagares.Pagare.Persona[5].Nombre2 + ' ' + pagares.Pagare.Persona[5].Apellido1 + ' ' + pagares.Pagare.Persona[5].Apellido2,
                TIPODOCCOD2 = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                NUMIDCOD2 = pagares.Pagare.Persona[5].Documento.Numero,
                DIRCOD2 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[6].Ubicacion.Direccion) ? pagares.Pagare.Persona[6].Ubicacion.Direccion : pagares.Pagare.Persona[5].Ubicacion.Direccion,
                TELCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? (pagares.Pagare.Persona[6].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[6].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[5].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].Ubicacion.Ciudad : pagares.Pagare.Persona[5].Ubicacion.Ciudad,
                CORREOCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].DatosContacto.Email : pagares.Pagare.Persona[5].DatosContacto.Email,

                // Asignación de valores para persona[6]
                RAZONSOCCOD2 = pagares.Pagare.Persona[6].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD2 = pagares.Pagare.Persona[6].Documento.TipoDocumento.Id,
                NUMIDRSCOD2 = pagares.Pagare.Persona[6].Documento.Numero,

                // Asignación de valores para persona[7]
                NOMAPCOD3 = pagares.Pagare.Persona[7].Nombre1 + ' ' + pagares.Pagare.Persona[7].Nombre2 + ' ' + pagares.Pagare.Persona[7].Apellido1 + ' ' + pagares.Pagare.Persona[7].Apellido2,
                TIPODOCCOD3 = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                NUMIDCOD3 = pagares.Pagare.Persona[7].Documento.Numero,
                DIRCOD3 = pagares.Pagare.Persona[7].Ubicacion.Direccion,
                TELCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? (pagares.Pagare.Persona[8].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[8].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[7].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].Ubicacion.Ciudad : pagares.Pagare.Persona[7].Ubicacion.Ciudad,
                CORREOCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].DatosContacto.Email : pagares.Pagare.Persona[7].DatosContacto.Email,

                // Asignación de valores para persona[7]
                RAZONSOCCOD3 = pagares.Pagare.Persona[8].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD3 = pagares.Pagare.Persona[8].Documento.TipoDocumento.Id,
                NUMIDRSCOD3 = pagares.Pagare.Persona[8].Documento.Numero,

                // Asignación de valores para persona[9]
                NOMAPINT1 = pagares.Pagare.Persona[9].Nombre1 + ' ' + pagares.Pagare.Persona[9].Nombre2 + ' ' + pagares.Pagare.Persona[9].Apellido1 + ' ' + pagares.Pagare.Persona[9].Apellido2,
                TIPODOCINT1 = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                NUMIDINT1 = pagares.Pagare.Persona[9].Documento.Numero,
                DIRINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].Ubicacion.Direccion : pagares.Pagare.Persona[9].Ubicacion.Direccion,
                TELINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? (pagares.Pagare.Persona[10].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[10].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[9].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].Ubicacion.Ciudad : pagares.Pagare.Persona[9].Ubicacion.Ciudad,
                CORREOINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].DatosContacto.Email : pagares.Pagare.Persona[9].DatosContacto.Email,

                // Asignación de valores para persona[10]
                RAZONSOCINT1 = pagares.Pagare.Persona[10].Nombre1 + ' ' + pagares.Pagare.Persona[10].Nombre2 + ' ' + pagares.Pagare.Persona[10].Apellido1 + ' ' + pagares.Pagare.Persona[10].Apellido2,
                TIPODOCCONUNINT1 = pagares.Pagare.Persona[10].Documento.TipoDocumento.Id,
                NUMIDCONUNINT1 = pagares.Pagare.Persona[10].Documento.Numero,

                // Asignación de valores para persona[11]
                NOMAPINT2 = pagares.Pagare.Persona[11].Nombre1 + ' ' + pagares.Pagare.Persona[11].Nombre2 + ' ' + pagares.Pagare.Persona[11].Apellido1 + ' ' + pagares.Pagare.Persona[11].Apellido2,
                TIPODOCINT2 = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                NUMIDINT2 = pagares.Pagare.Persona[11].Documento.Numero,
                DIRINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].Ubicacion.Direccion : pagares.Pagare.Persona[11].Ubicacion.Direccion,
                TELINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? (pagares.Pagare.Persona[12].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[12].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[11].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].Ubicacion.Ciudad : pagares.Pagare.Persona[11].Ubicacion.Ciudad,
                CORREOINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].DatosContacto.Email : pagares.Pagare.Persona[11].DatosContacto.Email,

                // Asignación de valores para persona[12]
                RAZONSOCINT2 = pagares.Pagare.Persona[12].Nombre1 + ' ' + pagares.Pagare.Persona[12].Nombre2 + ' ' + pagares.Pagare.Persona[12].Apellido1 + ' ' + pagares.Pagare.Persona[12].Apellido2,
                TIPODOCCONUNINT2 = pagares.Pagare.Persona[12].Documento.TipoDocumento.Id,
                NUMIDCONUNINT2 = pagares.Pagare.Persona[12].Documento.Numero,

                // Asignación de valores para persona[13]
                NOMAPINT3 = pagares.Pagare.Persona[13].Nombre1 + ' ' + pagares.Pagare.Persona[13].Nombre2 + ' ' + pagares.Pagare.Persona[13].Apellido1 + ' ' + pagares.Pagare.Persona[13].Apellido2,
                TIPODOCINT3 = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                NUMIDINT3 = pagares.Pagare.Persona[13].Documento.Numero,
                DIRINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].Ubicacion.Direccion : pagares.Pagare.Persona[13].Ubicacion.Direccion,
                TELINT3 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? (pagares.Pagare.Persona[14].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[14].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[13].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].Ubicacion.Ciudad : pagares.Pagare.Persona[13].Ubicacion.Ciudad,
                CORREOINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].DatosContacto.Email : pagares.Pagare.Persona[13].DatosContacto.Email,

                // Asignación de valores para persona[14]
                RAZONSOCINT3 = pagares.Pagare.Persona[14].Nombre1 + ' ' + pagares.Pagare.Persona[14].Nombre2 + ' ' + pagares.Pagare.Persona[14].Apellido1 + ' ' + pagares.Pagare.Persona[14].Apellido2,
                TIPODOCCONUNINT3 = pagares.Pagare.Persona[14].Documento.TipoDocumento.Id,
                NUMIDCONUNINT3 = pagares.Pagare.Persona[14].Documento.Numero,


            };


            documents = new List<Documents>
                    {
                        new Documents
                         {

                                TemplateId = TipoPagare,
                                TemplateFields = templateFields,
                                DocumentName = pagares.Pagare.Id,
                                Main = true,
                                SignatureRequired = true,




                         },
                        new Documents
                         {

                                TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                Main = false,
                                SignatureRequired = true,




                         },


                    };
            break;
        case 20:
        case 22:
        case 23:
            signers = new List<Signer>();
            if (pagares.Pagare.Persona.Count > 1 && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[1].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[1].Nombre1,
                    LastName = pagares.Pagare.Persona[1].Apellido1,
                    DocumentId = pagares.Pagare.Persona[1].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 9 && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[9].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[9].Nombre1,
                    LastName = pagares.Pagare.Persona[9].Apellido1,
                    DocumentId = pagares.Pagare.Persona[9].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 11 && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[11].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[11].Nombre1,
                    LastName = pagares.Pagare.Persona[11].Apellido1,
                    DocumentId = pagares.Pagare.Persona[11].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 13 && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[13].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[13].Nombre1,
                    LastName = pagares.Pagare.Persona[13].Apellido1,
                    DocumentId = pagares.Pagare.Persona[13].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 15 && !string.IsNullOrEmpty(pagares.Pagare.Persona[15].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[15].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[15].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[15].Nombre1,
                    LastName = pagares.Pagare.Persona[15].Apellido1,
                    DocumentId = pagares.Pagare.Persona[15].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[15].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[15].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 17 && !string.IsNullOrEmpty(pagares.Pagare.Persona[17].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[17].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[17].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[17].Nombre1,
                    LastName = pagares.Pagare.Persona[17].Apellido1,
                    DocumentId = pagares.Pagare.Persona[17].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[17].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[17].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }

            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.NumeroCelular.ToString() : pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].Ubicacion.Ciudad : pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.Email : pagares.Pagare.Persona[1].DatosContacto.Email,
                RAZONSOCTOM = pagares.Pagare.Persona[2].Nombre1 + ' ' + pagares.Pagare.Persona[2].Nombre2 + ' ' + pagares.Pagare.Persona[2].Apellido1 + ' ' + pagares.Pagare.Persona[2].Apellido2,
                TIPODOCTOMPJ = pagares.Pagare.Persona[2].Documento.TipoDocumento.Id,
                NUMIDTOMPJ = pagares.Pagare.Persona[2].Documento.Numero,


                // Asignación de valores para persona[9]
                NOMAPINT1 = pagares.Pagare.Persona[9].Nombre1 + ' ' + pagares.Pagare.Persona[9].Nombre2 + ' ' + pagares.Pagare.Persona[9].Apellido1 + ' ' + pagares.Pagare.Persona[9].Apellido2,
                TIPODOCINT1 = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                NUMIDINT1 = pagares.Pagare.Persona[9].Documento.Numero,
                DIRINT1 = pagares.Pagare.Persona[9].Ubicacion.Direccion,
                TELINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? (pagares.Pagare.Persona[10].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[10].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[9].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].Ubicacion.Ciudad : pagares.Pagare.Persona[9].Ubicacion.Ciudad,
                CORREOINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].DatosContacto.Email : pagares.Pagare.Persona[9].DatosContacto.Email,
                // Asignación de valores para persona[10]

                RAZONSOCINT1 = pagares.Pagare.Persona[10].Nombre1 + ' ' + pagares.Pagare.Persona[10].Nombre2 + ' ' + pagares.Pagare.Persona[10].Apellido1 + ' ' + pagares.Pagare.Persona[10].Apellido2,
                TIPODOCCONUNINT1 = pagares.Pagare.Persona[10].Documento.TipoDocumento.Id,
                NUMIDCONUNINT1 = pagares.Pagare.Persona[10].Documento.Numero,

                // Asignación de valores para persona[11]
                NOMAPINT2 = pagares.Pagare.Persona[11].Nombre1 + ' ' + pagares.Pagare.Persona[11].Nombre2 + ' ' + pagares.Pagare.Persona[11].Apellido1 + ' ' + pagares.Pagare.Persona[11].Apellido2,
                TIPODOCINT2 = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                NUMIDINT2 = pagares.Pagare.Persona[11].Documento.Numero,
                DIRINT2 = pagares.Pagare.Persona[11].Ubicacion.Direccion,
                TELINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? (pagares.Pagare.Persona[12].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[12].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[11].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].Ubicacion.Ciudad : pagares.Pagare.Persona[11].Ubicacion.Ciudad,
                CORREOINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].DatosContacto.Email : pagares.Pagare.Persona[11].DatosContacto.Email,

                // Asignación de valores para persona[12]
                RAZONSOCINT2 = pagares.Pagare.Persona[12].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[12].Apellido1 + ' ' + pagares.Pagare.Persona[12].Apellido2,
                TIPODOCCONUNINT2 = pagares.Pagare.Persona[12].Documento.TipoDocumento.Id,
                NUMIDCONUNINT2 = pagares.Pagare.Persona[12].Documento.Numero,

                // Asignación de valores para persona[13]
                NOMAPINT3 = pagares.Pagare.Persona[13].Nombre1 + ' ' + pagares.Pagare.Persona[13].Nombre2 + ' ' + pagares.Pagare.Persona[13].Apellido1 + ' ' + pagares.Pagare.Persona[13].Apellido2,
                TIPODOCINT3 = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                NUMIDINT3 = pagares.Pagare.Persona[13].Documento.Numero,
                DIRINT3 = pagares.Pagare.Persona[13].Ubicacion.Direccion,
                TELINT3 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? (pagares.Pagare.Persona[14].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[14].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[13].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].Ubicacion.Ciudad : pagares.Pagare.Persona[13].Ubicacion.Ciudad,
                CORREOINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].DatosContacto.Email : pagares.Pagare.Persona[13].DatosContacto.Email,

                // Asignación de valores para persona[14]
                RAZONSOCINT3 = pagares.Pagare.Persona[14].Nombre1 + ' ' + pagares.Pagare.Persona[13].Nombre2 + ' ' + pagares.Pagare.Persona[13].Apellido1 + ' ' + pagares.Pagare.Persona[13].Apellido2,
                TIPODOCCONUNINT3 = pagares.Pagare.Persona[14].Documento.TipoDocumento.Id,
                NUMIDCONUNINT3 = pagares.Pagare.Persona[14].Documento.Numero,

                // Asignación de valores para persona[15]
                NOMAPINT4 = pagares.Pagare.Persona[15].Nombre1 + ' ' + pagares.Pagare.Persona[15].Nombre2 + ' ' + pagares.Pagare.Persona[15].Apellido1 + ' ' + pagares.Pagare.Persona[15].Apellido2,
                TIPODOCINT4 = pagares.Pagare.Persona[15].Documento.TipoDocumento.Id,
                NUMIDINT4 = pagares.Pagare.Persona[15].Documento.Numero,
                DIRINT4 = !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? pagares.Pagare.Persona[16].Ubicacion.Direccion : pagares.Pagare.Persona[15].Ubicacion.Direccion,
                TELINT4 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? (pagares.Pagare.Persona[16].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[16].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[15].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[15].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT4 = !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? pagares.Pagare.Persona[16].Ubicacion.Ciudad : pagares.Pagare.Persona[15].Ubicacion.Ciudad,
                CORREOINT4 = !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? pagares.Pagare.Persona[16].DatosContacto.Email : pagares.Pagare.Persona[15].DatosContacto.Email,

                // Asignación de valores para persona[15]
                RAZONSOCINT4 = pagares.Pagare.Persona[15].Nombre1 + ' ' + pagares.Pagare.Persona[15].Nombre2 + ' ' + pagares.Pagare.Persona[15].Apellido1 + ' ' + pagares.Pagare.Persona[15].Apellido2,
                TIPODOCCONUNINT4 = pagares.Pagare.Persona[15].Documento.TipoDocumento.Id,
                NUMIDCONUNINT4 = pagares.Pagare.Persona[15].Documento.Numero,

                // Asignación de valores para persona[17]
                NOMAPINT5 = pagares.Pagare.Persona[17].Nombre1 + ' ' + pagares.Pagare.Persona[17].Nombre2 + ' ' + pagares.Pagare.Persona[17].Apellido1 + ' ' + pagares.Pagare.Persona[17].Apellido2,
                TIPODOCINT5 = pagares.Pagare.Persona[17].Documento.TipoDocumento.Id,
                NUMIDINT5 = pagares.Pagare.Persona[17].Documento.Numero,
                DIRINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? pagares.Pagare.Persona[18].Ubicacion.Direccion : pagares.Pagare.Persona[17].Ubicacion.Direccion,
                TELINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? (pagares.Pagare.Persona[18].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[18].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[17].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[17].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? pagares.Pagare.Persona[18].Ubicacion.Ciudad : pagares.Pagare.Persona[17].Ubicacion.Ciudad,
                CORREOINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? pagares.Pagare.Persona[18].DatosContacto.Email : pagares.Pagare.Persona[17].DatosContacto.Email,

                // Asignación de valores para persona[18]
                RAZONSOCINT5 = pagares.Pagare.Persona[18].Nombre1 + ' ' + pagares.Pagare.Persona[18].Nombre2 + ' ' + pagares.Pagare.Persona[18].Apellido1 + ' ' + pagares.Pagare.Persona[18].Apellido2,
                TIPODOCCONUNINT5 = pagares.Pagare.Persona[18].Documento.TipoDocumento.Id,
                NUMIDCONUNINT5 = pagares.Pagare.Persona[18].Documento.Numero,



            };


            documents = new List<Documents>
                    {
                        new Documents
                         {

                                TemplateId = TipoPagare,
                                TemplateFields = templateFields,
                                DocumentName = pagares.Pagare.Id,
                                Main = true,
                                SignatureRequired = true,




                         },
                        new Documents
                         {

                                TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                Main = false,
                                SignatureRequired = true,




                         },


                    };
            break;
        case 21:
        case 24:
            signers = new List<Signer>();
            if (pagares.Pagare.Persona.Count > 1 && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[1].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[1].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[1].Nombre1,
                    LastName = pagares.Pagare.Persona[1].Apellido1,
                    DocumentId = pagares.Pagare.Persona[1].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 3 && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[3].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[3].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[3].Nombre1,
                    LastName = pagares.Pagare.Persona[3].Apellido1,
                    DocumentId = pagares.Pagare.Persona[3].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 5 && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[5].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[5].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[5].Nombre1,
                    LastName = pagares.Pagare.Persona[5].Apellido1,
                    DocumentId = pagares.Pagare.Persona[5].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 7 && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[7].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[7].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[7].Nombre1,
                    LastName = pagares.Pagare.Persona[7].Apellido1,
                    DocumentId = pagares.Pagare.Persona[7].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 9 && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[9].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[9].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[9].Nombre1,
                    LastName = pagares.Pagare.Persona[9].Apellido1,
                    DocumentId = pagares.Pagare.Persona[9].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 11 && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[11].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[11].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[11].Nombre1,
                    LastName = pagares.Pagare.Persona[11].Apellido1,
                    DocumentId = pagares.Pagare.Persona[11].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 13 && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[13].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[13].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[13].Nombre1,
                    LastName = pagares.Pagare.Persona[13].Apellido1,
                    DocumentId = pagares.Pagare.Persona[13].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 15 && !string.IsNullOrEmpty(pagares.Pagare.Persona[15].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[15].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[15].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[15].Nombre1,
                    LastName = pagares.Pagare.Persona[15].Apellido1,
                    DocumentId = pagares.Pagare.Persona[15].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[15].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[15].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }
            if (pagares.Pagare.Persona.Count > 17 && !string.IsNullOrEmpty(pagares.Pagare.Persona[17].Documento.Numero) && !string.IsNullOrEmpty(pagares.Pagare.Persona[17].DatosContacto.Email))
            {
                signers.Add(new Signer
                {
                    ExternalUser = true,
                    Email = pagares.Pagare.Persona[17].DatosContacto.Email,
                    FirstName = pagares.Pagare.Persona[17].Nombre1,
                    LastName = pagares.Pagare.Persona[17].Apellido1,
                    DocumentId = pagares.Pagare.Persona[17].Documento.Numero,
                    DocumentType = pagares.Pagare.Persona[17].Documento.TipoDocumento.Id,
                    MobileNumber = pagares.Pagare.Persona[17].DatosContacto.NumeroCelular.ToString(),
                    CertificateProfile = "External Electronic"
                });
            }

            templateFields = new TemplateFields
            {
                CONSISE = pagares.Pagare.Numero,
                DIACREASISE = pagares.Pagare.FechaCreacion.Day.ToString(),
                MESCREASISE = pagares.Pagare.FechaCreacion.ToString("MMMM", new CultureInfo("es-ES")).ToUpper(),
                ANOCREASISE = pagares.Pagare.FechaCreacion.Year.ToString(),
                NOMAPTOM = pagares.Pagare.Persona[1].Nombre1 + ' ' + pagares.Pagare.Persona[1].Nombre2 + ' ' + pagares.Pagare.Persona[1].Apellido1 + ' ' + pagares.Pagare.Persona[1].Apellido2,
                TIPODOCTOMPN = pagares.Pagare.Persona[1].Documento.TipoDocumento.Id,
                NUMIDTOMPN = pagares.Pagare.Persona[1].Documento.Numero,
                DIRTOM = pagares.Pagare.Persona[1].Ubicacion.Direccion,
                TELTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.NumeroCelular.ToString() : pagares.Pagare.Persona[1].DatosContacto.NumeroCelular.ToString(),
                CIUDOMTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].Ubicacion.Ciudad : pagares.Pagare.Persona[1].Ubicacion.Ciudad,
                CORREOTOM = !string.IsNullOrEmpty(pagares.Pagare.Persona[2].Documento.Numero) ? pagares.Pagare.Persona[2].DatosContacto.Email : pagares.Pagare.Persona[1].DatosContacto.Email,
                RAZONSOCTOM = pagares.Pagare.Persona[2].Nombre1 + ' ' + pagares.Pagare.Persona[2].Nombre2 + ' ' + pagares.Pagare.Persona[2].Apellido1 + ' ' + pagares.Pagare.Persona[2].Apellido2,
                TIPODOCTOMPJ = pagares.Pagare.Persona[2].Documento.TipoDocumento.Id,
                NUMIDTOMPJ = pagares.Pagare.Persona[2].Documento.Numero,

                // Asignación de valores para persona[3]
                NOMAPCOD1 = pagares.Pagare.Persona[3].Nombre1 + ' ' + pagares.Pagare.Persona[3].Nombre2 + ' ' + pagares.Pagare.Persona[3].Apellido1 + ' ' + pagares.Pagare.Persona[3].Apellido2,
                TIPODOCCOD1 = pagares.Pagare.Persona[3].Documento.TipoDocumento.Id,
                NUMIDCOD1 = pagares.Pagare.Persona[3].Documento.Numero,
                DIRCOD1 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[4].Ubicacion.Direccion) ? pagares.Pagare.Persona[4].Ubicacion.Direccion : pagares.Pagare.Persona[3].Ubicacion.Direccion,
                TELCOD1  = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? (pagares.Pagare.Persona[4].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[4].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[3].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[3].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].Ubicacion.Ciudad : pagares.Pagare.Persona[3].Ubicacion.Ciudad,
                CORREOCOD1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[4].Documento.Numero) ? pagares.Pagare.Persona[4].DatosContacto.Email : pagares.Pagare.Persona[3].DatosContacto.Email,
                // Asignación de valores para persona[4]

                RAZONSOCCOD1 = pagares.Pagare.Persona[4].Nombre1 + ' ' + pagares.Pagare.Persona[4].Nombre2 + ' ' + pagares.Pagare.Persona[4].Apellido1 + ' ' + pagares.Pagare.Persona[4].Apellido2,
                TIPODOCRSCOD1 = pagares.Pagare.Persona[4].Documento.TipoDocumento.Id,
                NUMIDRSCOD1 = pagares.Pagare.Persona[4].Documento.Numero,

                // Asignación de valores para persona[5]
                NOMAPCOD2 = pagares.Pagare.Persona[5].Nombre1 + ' ' + pagares.Pagare.Persona[5].Nombre2 + ' ' + pagares.Pagare.Persona[5].Apellido1 + ' ' + pagares.Pagare.Persona[5].Apellido2,
                TIPODOCCOD2 = pagares.Pagare.Persona[5].Documento.TipoDocumento.Id,
                NUMIDCOD2 = pagares.Pagare.Persona[5].Documento.Numero,
                DIRCOD2 = !string.IsNullOrWhiteSpace(pagares.Pagare.Persona[6].Ubicacion.Direccion) ? pagares.Pagare.Persona[6].Ubicacion.Direccion : pagares.Pagare.Persona[5].Ubicacion.Direccion,
                TELCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? (pagares.Pagare.Persona[6].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[6].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[5].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[5].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].Ubicacion.Ciudad : pagares.Pagare.Persona[5].Ubicacion.Ciudad,
                CORREOCOD2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[6].Documento.Numero) ? pagares.Pagare.Persona[6].DatosContacto.Email : pagares.Pagare.Persona[5].DatosContacto.Email,

                // Asignación de valores para persona[6]
                RAZONSOCCOD2 = pagares.Pagare.Persona[6].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD2 = pagares.Pagare.Persona[6].Documento.TipoDocumento.Id,
                NUMIDRSCOD2 = pagares.Pagare.Persona[6].Documento.Numero,

                // Asignación de valores para persona[7]
                NOMAPCOD3 = pagares.Pagare.Persona[7].Nombre1 + ' ' + pagares.Pagare.Persona[7].Nombre2 + ' ' + pagares.Pagare.Persona[7].Apellido1 + ' ' + pagares.Pagare.Persona[7].Apellido2,
                TIPODOCCOD3 = pagares.Pagare.Persona[7].Documento.TipoDocumento.Id,
                NUMIDCOD3 = pagares.Pagare.Persona[7].Documento.Numero,
                DIRCOD3 = pagares.Pagare.Persona[7].Ubicacion.Direccion,
                TELCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? (pagares.Pagare.Persona[8].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[8].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[7].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[7].DatosContacto.NumeroCelular.ToString(),
                CIUDOMCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].Ubicacion.Ciudad : pagares.Pagare.Persona[7].Ubicacion.Ciudad,
                CORREOCOD3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[8].Documento.Numero) ? pagares.Pagare.Persona[8].DatosContacto.Email : pagares.Pagare.Persona[7].DatosContacto.Email,

                // Asignación de valores para persona[7]
                RAZONSOCCOD3 = pagares.Pagare.Persona[8].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[6].Apellido1 + ' ' + pagares.Pagare.Persona[6].Apellido2,
                TIPODOCRSCOD3 = pagares.Pagare.Persona[8].Documento.TipoDocumento.Id,
                NUMIDRSCOD3 = pagares.Pagare.Persona[8].Documento.Numero,

                // Asignación de valores para persona[9]
                NOMAPINT1 = pagares.Pagare.Persona[9].Nombre1 + ' ' + pagares.Pagare.Persona[9].Nombre2 + ' ' + pagares.Pagare.Persona[9].Apellido1 + ' ' + pagares.Pagare.Persona[9].Apellido2,
                TIPODOCINT1 = pagares.Pagare.Persona[9].Documento.TipoDocumento.Id,
                NUMIDINT1 = pagares.Pagare.Persona[9].Documento.Numero,
                DIRINT1 = pagares.Pagare.Persona[9].Ubicacion.Direccion,
                TELINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? (pagares.Pagare.Persona[10].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[10].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[9].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[9].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].Ubicacion.Ciudad : pagares.Pagare.Persona[9].Ubicacion.Ciudad,
                CORREOINT1 = !string.IsNullOrEmpty(pagares.Pagare.Persona[10].Documento.Numero) ? pagares.Pagare.Persona[10].DatosContacto.Email : pagares.Pagare.Persona[9].DatosContacto.Email,
                // Asignación de valores para persona[10]

                RAZONSOCINT1 = pagares.Pagare.Persona[10].Nombre1 + ' ' + pagares.Pagare.Persona[10].Nombre2 + ' ' + pagares.Pagare.Persona[10].Apellido1 + ' ' + pagares.Pagare.Persona[10].Apellido2,
                TIPODOCCONUNINT1 = pagares.Pagare.Persona[10].Documento.TipoDocumento.Id,
                NUMIDCONUNINT1 = pagares.Pagare.Persona[10].Documento.Numero,

                // Asignación de valores para persona[11]
                NOMAPINT2 = pagares.Pagare.Persona[11].Nombre1 + ' ' + pagares.Pagare.Persona[11].Nombre2 + ' ' + pagares.Pagare.Persona[11].Apellido1 + ' ' + pagares.Pagare.Persona[11].Apellido2,
                TIPODOCINT2 = pagares.Pagare.Persona[11].Documento.TipoDocumento.Id,
                NUMIDINT2 = pagares.Pagare.Persona[11].Documento.Numero,
                DIRINT2 = pagares.Pagare.Persona[11].Ubicacion.Direccion,
                TELINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? (pagares.Pagare.Persona[12].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[12].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[11].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[11].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].Ubicacion.Ciudad : pagares.Pagare.Persona[11].Ubicacion.Ciudad,
                CORREOINT2 = !string.IsNullOrEmpty(pagares.Pagare.Persona[12].Documento.Numero) ? pagares.Pagare.Persona[12].DatosContacto.Email : pagares.Pagare.Persona[11].DatosContacto.Email,

                // Asignación de valores para persona[12]
                RAZONSOCINT2 = pagares.Pagare.Persona[12].Nombre1 + ' ' + pagares.Pagare.Persona[6].Nombre2 + ' ' + pagares.Pagare.Persona[12].Apellido1 + ' ' + pagares.Pagare.Persona[12].Apellido2,
                TIPODOCCONUNINT2 = pagares.Pagare.Persona[12].Documento.TipoDocumento.Id,
                NUMIDCONUNINT2 = pagares.Pagare.Persona[12].Documento.Numero,

                // Asignación de valores para persona[13]
                NOMAPINT3 = pagares.Pagare.Persona[13].Nombre1 + ' ' + pagares.Pagare.Persona[13].Nombre2 + ' ' + pagares.Pagare.Persona[13].Apellido1 + ' ' + pagares.Pagare.Persona[13].Apellido2,
                TIPODOCINT3 = pagares.Pagare.Persona[13].Documento.TipoDocumento.Id,
                NUMIDINT3 = pagares.Pagare.Persona[13].Documento.Numero,
                DIRINT3 = pagares.Pagare.Persona[13].Ubicacion.Direccion,
                TELINT3 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? (pagares.Pagare.Persona[14].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[14].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[13].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[13].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].Ubicacion.Ciudad : pagares.Pagare.Persona[13].Ubicacion.Ciudad,
                CORREOINT3 = !string.IsNullOrEmpty(pagares.Pagare.Persona[14].Documento.Numero) ? pagares.Pagare.Persona[14].DatosContacto.Email : pagares.Pagare.Persona[13].DatosContacto.Email,

                // Asignación de valores para persona[14]
                RAZONSOCINT3 = pagares.Pagare.Persona[14].Nombre1 + ' ' + pagares.Pagare.Persona[13].Nombre2 + ' ' + pagares.Pagare.Persona[13].Apellido1 + ' ' + pagares.Pagare.Persona[13].Apellido2,
                TIPODOCCONUNINT3 = pagares.Pagare.Persona[14].Documento.TipoDocumento.Id,
                NUMIDCONUNINT3 = pagares.Pagare.Persona[14].Documento.Numero,

                // Asignación de valores para persona[15]
                NOMAPINT4 = pagares.Pagare.Persona[15].Nombre1 + ' ' + pagares.Pagare.Persona[15].Nombre2 + ' ' + pagares.Pagare.Persona[15].Apellido1 + ' ' + pagares.Pagare.Persona[15].Apellido2,
                TIPODOCINT4 = pagares.Pagare.Persona[15].Documento.TipoDocumento.Id,
                NUMIDINT4 = pagares.Pagare.Persona[15].Documento.Numero,
                DIRINT4 = !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? pagares.Pagare.Persona[16].Ubicacion.Direccion : pagares.Pagare.Persona[15].Ubicacion.Direccion,
                TELINT4 =  !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? (pagares.Pagare.Persona[16].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[16].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[15].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[15].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT4 = !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? pagares.Pagare.Persona[16].Ubicacion.Ciudad : pagares.Pagare.Persona[15].Ubicacion.Ciudad,
                CORREOINT4 = !string.IsNullOrEmpty(pagares.Pagare.Persona[16].Documento.Numero) ? pagares.Pagare.Persona[16].DatosContacto.Email : pagares.Pagare.Persona[15].DatosContacto.Email,

                // Asignación de valores para persona[15]
                RAZONSOCINT4 = pagares.Pagare.Persona[15].Nombre1 + ' ' + pagares.Pagare.Persona[15].Nombre2 + ' ' + pagares.Pagare.Persona[15].Apellido1 + ' ' + pagares.Pagare.Persona[15].Apellido2,
                TIPODOCCONUNINT4 = pagares.Pagare.Persona[15].Documento.TipoDocumento.Id,
                NUMIDCONUNINT4 = pagares.Pagare.Persona[15].Documento.Numero,

                // Asignación de valores para persona[17]
                NOMAPINT5 = pagares.Pagare.Persona[17].Nombre1 + ' ' + pagares.Pagare.Persona[17].Nombre2 + ' ' + pagares.Pagare.Persona[17].Apellido1 + ' ' + pagares.Pagare.Persona[17].Apellido2,
                TIPODOCINT5 = pagares.Pagare.Persona[17].Documento.TipoDocumento.Id,
                NUMIDINT5 = pagares.Pagare.Persona[17].Documento.Numero,
                DIRINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? pagares.Pagare.Persona[18].Ubicacion.Direccion : pagares.Pagare.Persona[17].Ubicacion.Direccion,
                TELINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? (pagares.Pagare.Persona[18].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[18].DatosContacto.NumeroCelular.ToString() : (pagares.Pagare.Persona[17].DatosContacto.NumeroCelular == 0)
                ? ""
                : pagares.Pagare.Persona[17].DatosContacto.NumeroCelular.ToString(),
                CIUDOMINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? pagares.Pagare.Persona[18].Ubicacion.Ciudad : pagares.Pagare.Persona[17].Ubicacion.Ciudad,
                CORREOINT5 = !string.IsNullOrEmpty(pagares.Pagare.Persona[18].Documento.Numero) ? pagares.Pagare.Persona[18].DatosContacto.Email : pagares.Pagare.Persona[17].DatosContacto.Email,

                // Asignación de valores para persona[18]
                RAZONSOCINT5 = pagares.Pagare.Persona[18].Nombre1 + ' ' + pagares.Pagare.Persona[18].Nombre2 + ' ' + pagares.Pagare.Persona[18].Apellido1 + ' ' + pagares.Pagare.Persona[18].Apellido2,
                TIPODOCCONUNINT5 = pagares.Pagare.Persona[18].Documento.TipoDocumento.Id,
                NUMIDCONUNINT5 = pagares.Pagare.Persona[18].Documento.Numero,



            };


            documents = new List<Documents>
                    {
                        new Documents
                         {

                                TemplateId = TipoPagare,
                                TemplateFields = templateFields,
                                DocumentName = pagares.Pagare.Id,
                                Main = true,
                                SignatureRequired = true,




                         },
                        new Documents
                         {

                                TemplateId = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                DocumentName = "Anexo 21 Politica Tratamiendo de Datos Personales.pdf",
                                Main = false,
                                SignatureRequired = true,




                         },


                    };
            break;



        default:
            // Lógica para cualquier otro valor no especificado
            break;
    }


    if (pagares?.Pagare?.FrecuenciaRecordatorios != null && pagares.Pagare.FrecuenciaRecordatorios > 0)
    {
        notificationfrequency = new notificationFrequency
        {
            hours = pagares.Pagare.FrecuenciaRecordatorios,
            times = pagares.Pagare.NumeroRecordatorios
        };
    }

    var creationDocumentRequest = new CreationDocumentRequest
    {
        Documents = documents,
        Sequential = false,
        ContentType = ContentType.PROMISE_NOTE,
        AccountSignature = false,
        notificationMessage = pagares.Pagare.MensajeNotificacion,
        notificationFrequency = notificationfrequency,
        OwnerEmail = Environment.GetEnvironmentVariable("Certicamara__Subscription_Owner"),
        expirationDate = pagares?.Pagare?.FechaCaducidad != DateTime.MinValue ? pagares.Pagare.FechaCaducidad.ToUniversalTime().AddHours(5).ToString("yyyy-MM-ddTHH:mm:ss.fffZ") : null,
        signatureDeadline = pagares.Pagare.HorasBloqueo,
        Signers = signers,
        TransactionCode = pagares.Pagare.Transaccion.ControlId
    };

    var jsonPagares = JsonSerializer.Serialize(creationDocumentRequest, new JsonSerializerOptions { WriteIndented = true });
    return creationDocumentRequest;
}
        #endregion
    }